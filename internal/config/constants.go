// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/check-illiad
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

const (

	// MyAppName is the public name of this application.
	myAppName string = "check-illiad"

	// MyAppURL is the location of the repo for this application.
	myAppURL string = "https://github.com/atc0005/check-illiad"

	// MyAppDescription is the description for this application shown in
	// HelpText output.
	myAppDescription string = "Go-based tooling to check/verify an ILLiad server instance"
)

// Default (flag, config file, etc) settings if not overridden by user input.
const (
	defaultLogLevel            string = "info"
	defaultEmitBranding        bool   = false
	defaultIgnoreMissingEmails bool   = false
	defaultDBServerHost        string = ""
	defaultDBServerInstance    string = ""
	defaultDBServerPort        int    = 1433
	defaultDBServerUsername    string = ""
	defaultDBServerPassword    string = ""
	defaultDBServerTrustCert   bool   = false
	defaultDBName              string = ILLiadDatabase

	// Default threshold values for pending email notifications. These initial
	// pre-v0.1.0 values are somewhat arbitrary, but based on our existing
	// mail queue (local) and inbox (remote) monitoring thresholds.
	defaultThresholdCountWarning  int = 1  // probably normal, but worth noting
	defaultThresholdCountCritical int = 3  // a spike, worth a closer look
	defaultThresholdAgeWarning    int = 5  // at 5 minutes something is stuck
	defaultThresholdAgeCritical   int = 10 // at 10 minutes it is no longer a blip

	// defaultDBServerEncryptMode is a string value representing one of
	// `disable`, `false` or `true`. These values directly match the valid
	// values for the `denisenkom/go-mssqldb` driver parameter named
	// `encrypt`.
	defaultDBServerEncryptMode string = "false"
)

// These values directly map to the `encrypt` parameter settings used by the
// database driver.
// https://github.com/denisenkom/go-mssqldb#common-parameters
const (
	EncryptModeDisable string = "disable"
	EncryptModeTrue    string = "true"
	EncryptModeFalse   string = "false"
)

// https://docs.microsoft.com/en-us/previous-versions/sql/sql-server-2008-r2/ms143531(v=sql.105)
// https://docs.microsoft.com/en-us/sql/relational-databases/system-stored-procedures/sp-server-info-transact-sql
// https://docs.microsoft.com/en-us/sql/t-sql/statements/create-database-transact-sql
// https://docs.microsoft.com/en-us/sql/t-sql/statements/create-user-transact-sql
const (
	MSSQLInstanceNameMaxChars int = 16
	MSSQLUsernameMaxChars     int = 128
	MSSQLPasswordMaxChars     int = 128
	MSSQLDatabaseNameMaxChars int = 123
)

// TCP port ranges
// http://www.iana.org/assignments/port-numbers
// Port numbers are assigned in various ways, based on three ranges: System
// Ports (0-1023), User Ports (1024-49151), and the Dynamic and/or Private
// Ports (49152-65535)
const (
	TCPReservedPort            int = 0
	TCPSystemPortStart         int = 1
	TCPSystemPortEnd           int = 1023
	TCPUserPortStart           int = 1024
	TCPUserPortEnd             int = 49151
	TCPDynamicPrivatePortStart int = 49152
	TCPDynamicPrivatePortEnd   int = 65535
)

const (
	// ILLiadDatabase is the default database name for ILLiad software.
	ILLiadDatabase string = "ILLData"

	// ILLiadDatabaseEMailTable is the name of the database table which holds
	// metadata for email notifications.
	ILLiadDatabaseEMailTable string = "EMailCopies"
)

// Queries used to interact with the ILLiad database that this plugin is
// responsible for checking.
const (
	QueryILLiadEMailCancelledCount       string = "SELECT COUNT(*) As [CancelledEmails] FROM [ILLData].[dbo].[EMailCopies] WHERE [Status] = 'Cancelled';"
	QueryILLiadEMailNULLCount            string = "SELECT COUNT(*) As [NULLEmails] FROM [ILLData].[dbo].[EMailCopies] WHERE [Status] IS NULL;"
	QueryILLiadEMailSentCount            string = "SELECT COUNT(*) As [SentEmails] FROM [ILLData].[dbo].[EMailCopies] WHERE [Status] = 'Sent';"
	QueryILLiadEMailPendingCount         string = "SELECT COUNT(*) As [PendingEmails] FROM [ILLData].[dbo].[EMailCopies] WHERE [Status] = 'Pending';"
	QueryILLiadEMailStatusSummary        string = "SELECT [Status], COUNT(*) AS [Count] FROM [ILLData].[dbo].[EMailCopies] GROUP BY [Status];"
	QueryILLiadEMailCancelledEmailValues string = "SELECT [TransactionNumber], [EMailDate], [Status], [Note] FROM [ILLData].[dbo].[EMailCopies] WHERE [Status] = 'Cancelled';"
	QueryILLiadEMailPendingEmailValues   string = "SELECT [TransactionNumber], [EMailDate], [Status], [Note] FROM [ILLData].[dbo].[EMailCopies] WHERE [Status] = 'Pending';"
)
