// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/check-illiad
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import "time"

// LogLevel returns the user-provided logging level or the default value if
// not provided.
func (c Config) LogLevel() string {

	switch {
	case c.Logging.Level != nil:
		return *c.Logging.Level
	default:
		return defaultLogLevel
	}
}

// EmitBranding returns the user-provided choice of whether branded output is
// emitted with check results or the default value if not provided.
func (c Config) EmitBranding() bool {
	switch {
	case c.Logging.EmitBranding != nil:
		return *c.Logging.EmitBranding
	default:
		return defaultEmitBranding
	}
}

// IgnoreMissingEmails returns the user-provided choice of whether missing
// email notification entries in the database should be treated as an OK state
// or the default value if not provided.
func (c Config) IgnoreMissingEmails() bool {
	switch {
	case c.Filters.IgnoreMissingEmails != nil:
		return *c.Filters.IgnoreMissingEmails
	default:
		return defaultIgnoreMissingEmails
	}
}

// DBServerHost returns the user-provided database server host or the default
// value if not provided.
func (c Config) DBServerHost() string {
	switch {
	case c.DBServer.Host != nil:
		return *c.DBServer.Host
	default:
		return defaultDBServerHost
	}
}

// DBServerPort returns the user-provided database server port or the default
// value if not provided.
func (c Config) DBServerPort() int {
	switch {
	case c.DBServer.Port != nil:
		return *c.DBServer.Port
	default:
		return defaultDBServerPort
	}
}

// DBServerInstance returns the user-provided database server instance or the
// default value if not provided.
func (c Config) DBServerInstance() string {
	switch {
	case c.DBServer.Instance != nil:
		return *c.DBServer.Instance
	default:
		return defaultDBServerInstance
	}
}

// DBServerUsername returns the user-provided database server username or the
// default value if not provided.
func (c Config) DBServerUsername() string {
	switch {
	case c.DBServer.Username != nil:
		return *c.DBServer.Username
	default:
		return defaultDBServerUsername
	}
}

// DBServerPassword returns the user-provided database server password or the
// default value if not provided.
func (c Config) DBServerPassword() string {
	switch {
	case c.DBServer.Password != nil:
		return *c.DBServer.Password
	default:
		return defaultDBServerPassword
	}
}

// DBServerEncryptMode returns the user-provided encrypt mode or the default
// value if not provided.
func (c Config) DBServerEncryptMode() string {
	switch {
	case c.DBServer.EncryptMode != nil:
		return *c.DBServer.EncryptMode
	default:
		return defaultDBServerEncryptMode
	}
}

// DBServerTrustCert returns the user-provided choice of whether the database
// server certificate is trusted as-is or if validation is enforced, or the
// default value if not provided.
func (c Config) DBServerTrustCert() bool {
	switch {
	case c.DBServer.TrustCert != nil:
		return *c.DBServer.TrustCert
	default:
		return defaultDBServerTrustCert
	}
}

// DBName returns the user-provided database name or the default value if not
// provided.
func (c Config) DBName() string {
	switch {
	case c.Database.Name != nil:
		return *c.Database.Name
	default:
		return defaultDBName
	}
}

// EmailCountThresholds returns the user-provided WARNING and CRITICAL
// threshold values for the number of email notifications in a pending state,
// or the default threshold values if not provided.
func (c Config) EmailCountThresholds() EmailCountThresholds {
	switch {
	case c.Thresholds.CountWarning != nil && c.Thresholds.CountCritical != nil:
		return EmailCountThresholds{
			Critical: *c.Thresholds.CountCritical,
			Warning:  *c.Thresholds.CountWarning,
			Set:      true,
		}
	default:
		return EmailCountThresholds{
			Critical: defaultThresholdCountCritical,
			Warning:  defaultThresholdCountWarning,
			Set:      false,
		}
	}
}

// EmailAgeThresholds returns the user-provided WARNING and CRITICAL threshold
// values for the age of email notifications in a pending state, or the
// default threshold values if not provided.
func (c Config) EmailAgeThresholds() EmailAgeThresholds {
	switch {
	case c.Thresholds.CountWarning != nil && c.Thresholds.CountCritical != nil:
		return EmailAgeThresholds{
			Critical: time.Minute * time.Duration(*c.Thresholds.AgeCritical),
			Warning:  time.Minute * time.Duration(*c.Thresholds.AgeWarning),
			Set:      true,
		}
	default:
		return EmailAgeThresholds{
			Critical: time.Minute * time.Duration(defaultThresholdAgeCritical),
			Warning:  time.Minute * time.Duration(defaultThresholdAgeWarning),
			Set:      false,
		}
	}
}

// func (c Config) EmailAgeThresholds() EmailAgeThresholds {
// 	now := time.Now()
// 	switch {
// 	case c.Thresholds.CountWarning != nil && c.Thresholds.CountCritical != nil:
//
// 		return EmailAgeThresholds{
// 			Critical: now.Add(-1 * (time.Minute * time.Duration(*c.Thresholds.AgeCritical))),
// 			Warning:  now.Add(-1 * (time.Minute * time.Duration(*c.Thresholds.AgeWarning))),
// 			Set:      true,
// 		}
// 	default:
// 		return EmailAgeThresholds{
// 			Critical: now.Add(-1 * (time.Minute * time.Duration(defaultThresholdAgeCritical))),
// 			Warning:  now.Add(-1 * (time.Minute * time.Duration(defaultThresholdAgeWarning))),
// 			Set:      false,
// 		}
// 	}
// }
