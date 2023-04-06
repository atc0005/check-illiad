// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/check-illiad
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/rs/zerolog/log"

	"github.com/atc0005/check-illiad/internal/config"
	"github.com/atc0005/check-illiad/internal/dbqs"
	"github.com/atc0005/go-nagios"
)

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

func main() {

	plugin := nagios.NewPlugin()

	// defer this from the start so it is the last deferred function to run
	defer plugin.ReturnCheckResults()

	cfg, configErr := config.New()
	if configErr != nil {
		log.Err(configErr).Msg("error validating configuration")

		plugin.AddError(configErr)
		plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
		plugin.ServiceOutput = fmt.Sprintf(
			"%s: Failed to load configuration: %v",
			nagios.StateUNKNOWNLabel,
			configErr,
		)

		// no need to go any further, we *want* to exit right away; we don't
		// have a working configuration and there isn't anything further to do
		return
	}

	// Flesh out plugin with some additional common details now that
	// the flags have been parsed.
	plugin.LongServiceOutput = fmt.Sprintf(
		"* Database Connection%s"+
			"** Host: %q%s"+
			"** Port: %d%s"+
			"** Instance: %q%s"+
			"** Database: %q%s"+
			"** Encryption Mode: %q%s"+
			"** Trust Server Cert (disable verification): %v%s",
		nagios.CheckOutputEOL,
		cfg.DBServerHost(), nagios.CheckOutputEOL,
		cfg.DBServerPort(), nagios.CheckOutputEOL,
		cfg.DBServerInstance(), nagios.CheckOutputEOL,
		cfg.DBName(), nagios.CheckOutputEOL,
		cfg.DBServerEncryptMode(), nagios.CheckOutputEOL,
		cfg.DBServerTrustCert(), nagios.CheckOutputEOL,
	)

	plugin.CriticalThreshold = fmt.Sprintf(
		"[Age: %v, Count: %v]",
		cfg.EmailAgeThresholds().Critical,
		cfg.EmailCountThresholds().Critical,
	)
	plugin.WarningThreshold = fmt.Sprintf(
		"[Age: %v, Count: %v]",
		cfg.EmailAgeThresholds().Warning,
		cfg.EmailCountThresholds().Warning,
	)

	// If enabled, show application details at end of notification
	if cfg.EmitBranding() {
		plugin.BrandingCallback = config.Branding("Notification generated by ")
	}

	// Use ADO connection string format. This format does not appear to
	// require character encoding in order to handle special characters (e.g., encoding `{` as `%7B`)
	// https://github.com/denisenkom/go-mssqldb#common-parameters
	// https://github.com/denisenkom/go-mssqldb#less-common-parameters
	dsn := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s;encrypt=%s;TrustServerCertificate=%t;app name=%s;",
		cfg.DBServerHost(),
		cfg.DBServerUsername(),
		cfg.DBServerPassword(),
		cfg.DBServerPort(),
		cfg.DBName(),
		cfg.DBServerEncryptMode(),
		cfg.DBServerTrustCert(),
		config.Version(),
	)

	if cfg.DBServerInstance() != "" {
		dsn += fmt.Sprintf("instance=%s;", cfg.DBServerInstance())
	}

	// Setup connection (which is established lazily)
	db, dbOpenErr := sql.Open("mssql", dsn)
	if dbOpenErr != nil {
		cfg.Log.Error().Err(dbOpenErr).Msg("open connection failed")

		plugin.AddError(dbOpenErr)
		plugin.ServiceOutput = fmt.Sprintf(
			"%s: Error connecting to %s",
			nagios.StateUNKNOWNLabel,
			cfg.DBServerHost(),
		)
		plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode

		// no need to go any further, we *want* to exit right away; we don't
		// have a connection to the remote server and there isn't anything
		// further we can do
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			cfg.Log.Error().Err(err).Msg("error closing connection to server")
		} else {
			cfg.Log.Debug().Msg("successfully closed connection to server")
		}
	}()

	// Use Ping() to create a connection and check for any errors
	if err := db.Ping(); err != nil {
		cfg.Log.Error().Err(err).Msg("error verifying connection to server")

		plugin.AddError(err)
		plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
		plugin.ServiceOutput = fmt.Sprintf(
			"%s: Failed to establish connection to database: %v",
			nagios.StateUNKNOWNLabel,
			err,
		)

		return
	}

	cfg.Log.Debug().
		Str("host", cfg.DBServerHost()).
		Str("instance", cfg.DBServerInstance()).
		Int("port", cfg.DBServerPort()).
		Str("database", cfg.DBName()).
		Str("username", cfg.DBServerUsername()).
		Str("encrypt_mode", cfg.DBServerEncryptMode()).
		Bool("validate_cert", !cfg.DBServerTrustCert()).
		Msg("connection established")

	// ensure there are rows in the source database table
	rowCount, rowCountErr := dbqs.RowsCount(db, config.ILLiadDatabaseEMailTable)
	if rowCountErr != nil {
		cfg.Log.Error().
			Err(rowCountErr).
			Str("table_name", config.ILLiadDatabaseEMailTable).
			Msg("failed to execute query")

		plugin.AddError(rowCountErr)
		plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
		plugin.ServiceOutput = fmt.Sprintf(
			"%s: Failed to execute query against database: %v",
			nagios.StateUNKNOWNLabel,
			rowCountErr,
		)

		return
	}

	// Legitimate scenarios include a fresh ILLiad installation or a database
	// that has been recently "cleaned" of historical data. Because this is a
	// rare occurrence, sysadmin has to opt into ignoring this situation.
	if rowCount == 0 && cfg.IgnoreMissingEmails() {

		cfg.Log.Error().
			Err(rowCountErr).
			Str("table_name", config.ILLiadDatabaseEMailTable).
			Msg("no rows found in table")

		plugin.AddError(rowCountErr)
		plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
		plugin.ServiceOutput = fmt.Sprintf(
			"%s: Email notifications history missing in table %q in database %q",
			nagios.StateUNKNOWNLabel,
			config.ILLiadDatabaseEMailTable,
			config.ILLiadDatabase,
		)

		return
	}

	var pendingEMailCount int
	if err := db.QueryRow(config.QueryILLiadEMailPendingCount).Scan(&pendingEMailCount); err != nil {
		cfg.Log.Error().Err(err).Msg("failed to execute query")

		plugin.AddError(err)
		plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
		plugin.ServiceOutput = fmt.Sprintf(
			"%s: Failed to retrieve pending emails count: %v",
			nagios.StateUNKNOWNLabel,
			err,
		)

		return
	}

	//
	// Review query results, assert within provided ranges.
	//

	if pendingEMailCount > 0 {

		rows, queryErr := db.Query(config.QueryILLiadEMailPendingEmailValues)
		// rows, queryErr := db.Query(config.QueryILLiadEMailCancelledEmailValues)
		if queryErr != nil {
			plugin.AddError(queryErr)
			plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
			cfg.Log.Error().Err(queryErr).Msg("failed to execute query")

			plugin.ServiceOutput = fmt.Sprintf(
				"%s: Failed to retrieve pending email notifications: %v",
				nagios.StateUNKNOWNLabel,
				queryErr,
			)

			return
		}

		defer func() {
			if err := rows.Close(); err != nil {
				cfg.Log.Error().Err(err).Msg("error closing rows object")
			} else {
				cfg.Log.Debug().Msg("successfully closed rows object")
			}
		}()

		var emailNotifyEntry config.TableEMailCopies
		emailNotifyEntries := make([]config.TableEMailCopies, 0, 10)
		for rows.Next() {

			scanErr := rows.Scan(
				&emailNotifyEntry.TransactionNumber,
				&emailNotifyEntry.EMailDate,
				&emailNotifyEntry.Status,
				&emailNotifyEntry.Note,
			)
			if scanErr != nil {
				plugin.AddError(scanErr)
				plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
				cfg.Log.Error().Err(scanErr).Msg("failed to retrieve email notification data")

				plugin.ServiceOutput = fmt.Sprintf(
					"%s: Failed to retrieve email notification data from table %q",
					nagios.StateUNKNOWNLabel,
					config.ILLiadDatabaseEMailTable,
				)

				return
			}
			cfg.Log.Debug().Msg("completed scanning from row object")

			// accumulate here; we will process separately
			emailNotifyEntries = append(emailNotifyEntries, emailNotifyEntry)

		}
		cfg.Log.Debug().Msg("completed scanning all row objects")

		if err := rows.Err(); err != nil {
			plugin.AddError(err)
			plugin.ExitStatusCode = nagios.StateUNKNOWNExitCode
			cfg.Log.Error().Err(err).Msg("error occurred during rows scan loop")

			plugin.ServiceOutput = fmt.Sprintf(
				"%s: Error occurred during email notification data retrieval from table %q",
				nagios.StateUNKNOWNLabel,
				config.ILLiadDatabaseEMailTable,
			)

			return
		}
		cfg.Log.Debug().Msg("no row object scanning errors encountered")

		//
		// AGE checks
		//

		now := time.Now()
		criticalTime := now.Add(-1 * (cfg.EmailAgeThresholds().Critical))
		warningTime := now.Add(-1 * (cfg.EmailAgeThresholds().Warning))

		for _, entry := range emailNotifyEntries {
			if cfg.LogLevel() == config.LogLevelDebug {
				fmt.Printf("%+v\n", entry)
			}

			summaryMsgTmpl := "%s: pending email notification %d found older than %v threshold [total pending: %d]"
			errOldPendingNotifications := fmt.Errorf("old pending notifications found")

			if !entry.EMailDate.After(warningTime) {
				cfg.Log.Error().
					Err(errOldPendingNotifications).
					Int("transaction", entry.TransactionNumber).
					Msg("")
			}

			switch {
			case entry.EMailDate.Before(criticalTime) || entry.EMailDate.Equal(criticalTime):

				plugin.AddError(errOldPendingNotifications)
				plugin.ExitStatusCode = nagios.StateCRITICALExitCode

				plugin.ServiceOutput = fmt.Sprintf(
					summaryMsgTmpl,
					nagios.StateCRITICALLabel,
					entry.TransactionNumber,
					cfg.EmailAgeThresholds().Critical,
					pendingEMailCount,
				)

				return

			case entry.EMailDate.Before(warningTime) || entry.EMailDate.Equal(warningTime):

				plugin.AddError(errOldPendingNotifications)
				plugin.ExitStatusCode = nagios.StateWARNINGExitCode

				plugin.ServiceOutput = fmt.Sprintf(
					summaryMsgTmpl,
					nagios.StateWARNINGLabel,
					entry.TransactionNumber,
					cfg.EmailAgeThresholds().Warning,
					pendingEMailCount,
				)

				return

			case entry.EMailDate.After(warningTime):
				cfg.Log.
					Debug().
					Int("transaction", entry.TransactionNumber).
					Msg("young pending email notification")

			}

		}

		//
		// COUNT checks (less important than AGE checks, so after)
		//

		if len(emailNotifyEntries) >= cfg.EmailCountThresholds().Critical ||
			len(emailNotifyEntries) >= cfg.EmailCountThresholds().Warning {

			summaryMsg := "pending notifications greater than specified thresholds"

			switch {
			case len(emailNotifyEntries) >= cfg.EmailCountThresholds().Critical:
				plugin.ExitStatusCode = nagios.StateCRITICALExitCode
			case len(emailNotifyEntries) >= cfg.EmailCountThresholds().Warning:
				plugin.ExitStatusCode = nagios.StateWARNINGExitCode
			}

			plugin.AddError(fmt.Errorf(summaryMsg))
			plugin.ServiceOutput = fmt.Sprintf(
				"%s: %s [CRITICAL: %v, WARNING: %v, total pending: %d]",
				nagios.StateCRITICALLabel,
				summaryMsg,
				cfg.EmailCountThresholds().Critical,
				cfg.EmailCountThresholds().Warning,
				pendingEMailCount,
			)

			return

		}

	}

	//
	// ALL CLEAR
	//

	statusMsg := fmt.Sprintf(
		"%d email notifications in a pending state found without crossing specified thresholds",
		pendingEMailCount,
	)

	cfg.Log.Debug().
		Bool("age_thresholds_provided", cfg.EmailAgeThresholds().Set).
		Bool("count_thresholds_provided", cfg.EmailCountThresholds().Set).
		Msg(statusMsg)

	plugin.ServiceOutput = fmt.Sprintf(
		"%s: %s",
		nagios.StateOKLabel,
		statusMsg,
	)

	plugin.ExitStatusCode = nagios.StateOKExitCode

	// implied return, allow plugin.ReturnCheckResults() to run
	// return

}
