// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/check-illiad
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"fmt"
	"unicode"

	"github.com/atc0005/check-illiad/internal/caller"
	"github.com/atc0005/go-nagios"
)

// validate verifies that user-provided and/or default values are acceptable.
//
// In most cases (e.g., where a setting is required), getter methods are
// checked instead of directly referencing the config struct. getter methods
// pass user-provided values through without modification; if a user did not
// specify a value, the default value is passed through for validation.
func (c Config) validate() error {

	myFuncName := caller.GetFuncName()

	switch c.LogLevel() {
	case LogLevelDisabled:
	case LogLevelPanic:
	case LogLevelFatal:
	case LogLevelError:
	case LogLevelWarn:
	case LogLevelInfo:
	case LogLevelDebug:
	case LogLevelTrace:
	default:
		return fmt.Errorf(
			"%s: invalid log level provided: %v",
			myFuncName,
			c.LogLevel(),
		)
	}

	// EmitBranding returns a boolean value, so nothing to test here.

	if c.DBServerHost() == "" {
		return fmt.Errorf(
			"%s: missing database host name or IP Address",
			myFuncName,
		)
	}

	if !(c.DBServerPort() >= TCPSystemPortStart) && (c.DBServerPort() <= TCPDynamicPrivatePortEnd) {
		return fmt.Errorf(
			"%s: invalid port %d specified; outside of the range of %d and %d",
			myFuncName,
			c.DBServerPort(),
			TCPSystemPortStart,
			TCPDynamicPrivatePortEnd,
		)
	}

	// TODO: Extend this as needed to apply proper validation
	// https://docs.microsoft.com/en-us/previous-versions/sql/sql-server-2008-r2/ms143531(v=sql.105)
	// https://stackoverflow.com/questions/5260650/max-length-of-sql-server-instance-name
	switch {
	case c.DBServerInstance() == "":
		// TODO: How else to determine whether this was set without accessing
		// the field directly or returning a boolean value? Custom `Option`
		// struct type with fields for `High`, `Low`, `Set`?
		if c.DBServer.Instance != nil {
			return fmt.Errorf(
				"%s: missing database host instance",
				myFuncName,
			)
		}

	case !unicode.IsLetter(rune(c.DBServerInstance()[0])):
		return fmt.Errorf(
			"%s: specified instance name does not begin with a known Unicode letter; got %q, expected letter instead",
			myFuncName,
			c.DBServerInstance(),
		)

	case len(c.DBServerInstance()) > MSSQLInstanceNameMaxChars:
		return fmt.Errorf(
			"%s: specified instance name too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBServerInstance()),
			MSSQLInstanceNameMaxChars,
		)
	}

	switch {
	case c.DBServerUsername() == "":
		return fmt.Errorf(
			"%s: missing database host username",
			myFuncName,
		)
	case len(c.DBServerUsername()) > MSSQLUsernameMaxChars:
		return fmt.Errorf(
			"%s: specified username too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBServerUsername()),
			MSSQLUsernameMaxChars,
		)
	}

	switch {
	case c.DBServerPassword() == "":
		return fmt.Errorf(
			"%s: missing database host password",
			myFuncName,
		)
	case len(c.DBServerPassword()) > MSSQLPasswordMaxChars:
		return fmt.Errorf(
			"%s: specified password too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBServerPassword()),
			MSSQLPasswordMaxChars,
		)
	}

	switch c.DBServerEncryptMode() {
	case EncryptModeDisable:
	case EncryptModeTrue:
	case EncryptModeFalse:
	default:
		return fmt.Errorf(
			"%s: invalid option %q provided for connection encrypt mode",
			myFuncName,
			c.DBServerEncryptMode(),
		)
	}

	// DBServerTrustCert returns a boolean value, so nothing to test here.

	switch {
	case c.DBName() == "":
		return fmt.Errorf(
			"%s: missing database name",
			myFuncName,
		)

	case len(c.DBName()) > MSSQLDatabaseNameMaxChars:
		return fmt.Errorf(
			"%s: specified database name too long; got %d characters, max %d supported",
			myFuncName,
			len(c.DBName()),
			MSSQLDatabaseNameMaxChars,
		)

	}

	// If not specified, default threshold values are provided for WARNING and
	// CRITICAL states. We validate either way.
	emailCT := c.EmailCountThresholds()
	switch {
	case emailCT.Warning > emailCT.Critical:
		return fmt.Errorf(
			"%s: provided %s pending emails count (%v) greater than %s threshold (%v)",
			myFuncName,
			nagios.StateWARNINGLabel,
			emailCT.Warning,
			nagios.StateCRITICALLabel,
			emailCT.Critical,
		)

	case emailCT.Warning <= 0:
		return fmt.Errorf(
			"%s: invalid value (%v) provided for %s pending emails count",
			myFuncName,
			emailCT.Warning,
			nagios.StateWARNINGLabel,
		)

	case emailCT.Critical <= 0:
		return fmt.Errorf(
			"%s: invalid value (%v) provided for %s pending emails count",
			myFuncName,
			emailCT.Critical,
			nagios.StateCRITICALLabel,
		)

	}

	// If not specified, default threshold values are provided for WARNING and
	// CRITICAL states. We validate either way.
	emailAT := c.EmailAgeThresholds()
	switch {
	case emailAT.Warning > emailAT.Critical:
		return fmt.Errorf(
			"%s: provided %s pending emails age (%v) greater than %s threshold (%v)",
			myFuncName,
			nagios.StateWARNINGLabel,
			emailAT.Warning,
			nagios.StateCRITICALLabel,
			emailAT.Critical,
		)

	case emailAT.Warning <= 0:
		return fmt.Errorf(
			"%s: invalid value (%v) provided for %s pending emails age",
			myFuncName,
			emailAT.Warning,
			nagios.StateWARNINGLabel,
		)

	case emailAT.Critical <= 0:
		return fmt.Errorf(
			"%s: invalid value (%v) provided for %s pending emails age",
			myFuncName,
			emailAT.Critical,
			nagios.StateCRITICALLabel,
		)

	}

	// IgnoreMissingEmails returns a boolean value, so nothing to test here.

	return nil

}
