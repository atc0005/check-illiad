// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/check-illiad
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"database/sql"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/rs/zerolog"
)

// EmailAgeThresholds represents the user-specified email notification age
// thresholds.
type EmailAgeThresholds struct {
	Critical time.Duration `arg:"-"`
	Warning  time.Duration `arg:"-"`
	Set      bool          `arg:"-"`
}

// EmailCountThresholds represents the user-specified email notification count
// thresholds.
type EmailCountThresholds struct {
	Critical int  `arg:"-"`
	Warning  int  `arg:"-"`
	Set      bool `arg:"-"`
}

// DBServer is the user-specified settings for creating a database server
// connection.
type DBServer struct {
	Host        *string `arg:"--host,env:CHECK_ILLIAD_DBSERVER_HOST" help:"The hostname of the database server hosting the database used by the ILLiad software. If using encryption, this value should match one of the Subject Alternate Name (SANs) values listed on the certificate."`
	Port        *int    `arg:"--port,env:CHECK_ILLIAD_DBSERVER_PORT" help:"The TCP port used to connect to the database server. If not specified, the default port will be used."`
	Instance    *string `arg:"--instance,env:CHECK_ILLIAD_DBSERVER_INSTANCE" help:"The database server instance name. This may be blank."`
	Username    *string `arg:"--username,env:CHECK_ILLIAD_DBSERVER_USERNAME" help:"The username used to connect to the database server. An account with read-only access to the database used by the ILLiad software is sufficient."`
	Password    *string `arg:"--password,env:CHECK_ILLIAD_DBSERVER_PASSWORD" help:"The plaintext password used to connect to the database server. An account with read-only access to the database used by the ILLiad software is sufficient."`
	EncryptMode *string `arg:"--encrypt-mode,env:CHECK_ILLIAD_DBSERVER_ENCRYPT_MODE" help:"Whether data sent between client and server is encrypted. true for yes, false for login packet only and disable for no encryption."`
	TrustCert   *bool   `arg:"--trust-cert,env:CHECK_ILLIAD_TRUST_CERT" help:"Whether the certificate should be trusted as-is without validation. WARNING: TLS is susceptible to man-in-the-middle attacks if enabling this option."`
}

// TableEMailCopies reflects fields from the `EMailCopies` table from the
// `ILLData` database.
type TableEMailCopies struct {
	TransactionNumber int            `arg:"-"`
	EMailDate         time.Time      `arg:"-"`
	Status            sql.NullString `arg:"-"`
	Note              sql.NullString `arg:"-"`
}

// Database is the user-specified settings for the database used by the ILLiad
// software.
type Database struct {
	Name *string `arg:"--db-name,env:CHECK_ILLIAD_DATABASE_NAME" help:"The name of the database used by ILLiad software and checked by plugins from this project."`
}

// Logging represents options specific to how this application handles
// logging.
type Logging struct {
	Level        *string `arg:"--log-level,env:CHECK_ILLIAD_LOG_LEVEL" help:"Maximum log level at which messages will be logged. Log messages below this threshold will be discarded."`
	EmitBranding *bool   `arg:"--emit-branding,env:CHECK_ILLIAD_EMIT_BRANDING" help:"Whether 'generated by' text is included at the bottom of application output. This output is included in the Nagios dashboard and notifications. This output may not mix well with branding output from other tools such as atc0005/send2teams which also insert their own branding output."`
}

// Thresholds represents the values which determine WARNING and CRITICAL
// thresholds.
type Thresholds struct {
	CountWarning  *int `arg:"--count-warning,env:CHECK_ILLIAD_COUNT_WARNING" help:"The number of pending email notifications when this plugin will consider the service check to be in a WARNING state."`
	CountCritical *int `arg:"--count-critical,env:CHECK_ILLIAD_COUNT_CRITICAL" help:"The number of pending email notifications when this plugin will consider the service check to be in a CRITICAL state."`
	AgeWarning    *int `arg:"--age-warning,env:CHECK_ILLIAD_AGE_WARNING" help:"The number of minutes an email notification has been in a pending status when this plugin will consider the service check to be in a WARNING state."`
	AgeCritical   *int `arg:"--age-critical,env:CHECK_ILLIAD_AGE_CRITICAL" help:"The number of minutes an email notification has been in a pending status when this plugin will consider the service check to be in a CRITICAL state."`
}

// Filters represents options which lets the sysadmin filter in/out specific
// values or scenarios.
type Filters struct {
	IgnoreMissingEmails *bool `arg:"--ignore-missing-emails,env:CHECK_ILLIAD_IGNORE_MISSING_EMAILS" help:"Whether finding zero email notifications recorded in the database used by ILLiad software should be treated as an OK state. Legitimate scenarios include fresh ILLiad installations or recent purge of history."`
}

// Config is a unified set of configuration values for this application. This
// struct is configured via command-line flags or (maybe in the future) TOML
// configuration file provided by the user. Values held by this object are
// intended to be retrieved via "getter" methods.
type Config struct {
	Logging
	DBServer
	Database
	Thresholds
	Filters

	// Log is an embedded zerolog Logger initialized via config.New().
	Log zerolog.Logger `arg:"-"`

	flagParser *arg.Parser `arg:"-"`
}
