<!-- omit in toc -->
# check-illiad

Go-based tooling to check/verify an ILLiad server instance.

[![Latest Release](https://img.shields.io/github/release/atc0005/check-illiad.svg?style=flat-square)](https://github.com/atc0005/check-illiad/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/check-illiad.svg)](https://pkg.go.dev/github.com/atc0005/check-illiad)
[![Validate Codebase](https://github.com/atc0005/check-illiad/workflows/Validate%20Codebase/badge.svg)](https://github.com/atc0005/check-illiad/actions?query=workflow%3A%22Validate+Codebase%22)
[![Validate Docs](https://github.com/atc0005/check-illiad/workflows/Validate%20Docs/badge.svg)](https://github.com/atc0005/check-illiad/actions?query=workflow%3A%22Validate+Docs%22)
[![Lint and Build using Makefile](https://github.com/atc0005/check-illiad/workflows/Lint%20and%20Build%20using%20Makefile/badge.svg)](https://github.com/atc0005/check-illiad/actions?query=workflow%3A%22Lint+and+Build+using+Makefile%22)
[![Quick Validation](https://github.com/atc0005/check-illiad/workflows/Quick%20Validation/badge.svg)](https://github.com/atc0005/check-illiad/actions?query=workflow%3A%22Quick+Validation%22)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
- [Features](#features)
- [Status](#status)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
- [Configuration](#configuration)
  - [Precedence](#precedence)
  - [Command-line Arguments](#command-line-arguments)
  - [Environment Variables](#environment-variables)
- [Example](#example)
  - [Default thresholds](#default-thresholds)
- [License](#license)
- [Related projects](#related-projects)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

This repo is intended to provide Nagios plugins that may be used to monitor an
ILLiad instance. As of this writing, a plugin named `check_illiad_emails` is
available that may be used to monitor email notifications.

## Features

The `check_illiad_emails` plugin supports monitoring `Pending` email
notifications:

- greater number queued than specified
- that have remained "in the queue" for longer than a specified amount of time

Default values are provided, but are easily overridden with custom values in
order to match usage patterns for each instance.

## Status

As of this writing, this plugin is just leaving the bench for real-world use.
There may be bugs or minor issues present. Please report any issues that you
encounter.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go 1.14+
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Windows 7, Server 2008R2 or later
  - per official [Go install notes][go-docs-install]
- Windows 10 Version 1909
  - tested
- Ubuntu Linux 16.04, 18.04

## Installation

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
   - NOTE: Pay special attention to the remarks about `$HOME/.profile`
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/check-illiad`
   1. `cd check-illiad`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     - `sudo yum install make gcc`
   - for Windows
     - Emulated environments (*easier*)
       - Skip all of this and build using the default `go build` command in
         Windows (see below for use of the `-mod=vendor` flag)
       - build using Windows Subsystem for Linux Ubuntu environment and just
         copy out the Windows binaries from that environment
       - If already running a Docker environment, use a container with the Go
         tool-chain already installed
       - If already familiar with LXD, create a container and follow the
         installation steps given previously to install required dependencies
     - Native tooling (*harder*)
       - see the StackOverflow Question `32127524` link in the
         [References](references.md) section for potential options for
         installing `make` on Windows
       - see the mingw-w64 project homepage link in the
         [References](references.md) section for options for installing `gcc`
         and related packages on Windows
1. Build binaries
   - for the current operating system, explicitly using bundled dependencies
         in top-level `vendor` folder
     - `go build -mod=vendor ./cmd/check_illiad_emails/`
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for use on Windows
      - `make windows`
   - for use on Linux
     - `make linux`
1. Copy the newly compiled binary from the applicable `/tmp` subdirectory path
   (based on the clone instructions in this section) below and deploy where
   needed.
   - if using `Makefile`
     - look in `/tmp/check-illiad/release_assets/check_illiad_emails/`
   - if using `go build`
     - look in `/tmp/check-illiad/`

## Configuration

### Precedence

The priority order is:

1. Command line flags (highest priority)
1. Environment variables
1. Default settings (lowest priority)

In general, command-line options are the primary way of configuring settings
for this application, but environment variables are also a supported
alternative. Most plugin settings require that a value be specified by the
sysadmin, though some (e.g., logging) have useful defaults.

### Command-line Arguments

- Flags marked as **`required`** must be set via CLI flag or environment
  variable.
- Flags *not* marked as required are for settings where a useful default is
  already defined.

| Option                  | Required | Default        | Repeat | Possible                                                                | Description                                                                                                                                                                                                       |
| ----------------------- | -------- | -------------- | ------ | ----------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`             | No       | `false`        | No     | `h`, `help`                                                             | Show Help text along with the list of supported flags.                                                                                                                                                            |
| `emit-branding`         | No       | `false`        | No     | `true`, `false`                                                         | Toggles emission of branding details with plugin status details. This output is disabled by default.                                                                                                              |
| `log-level`             | No       | `info`         | No     | `disabled`, `panic`, `fatal`, `error`, `warn`, `info`, `debug`, `trace` | Log message priority filter. Log messages with a lower level are ignored.                                                                                                                                         |
| `host`                  | Yes      | *empty list*   | No     | *one or more valid files and directories*                               | The hostname of the database server hosting the database used by the ILLiad software. If using encryption, this value should match one of the Subject Alternate Name (SANs) values listed on the certificate.     |
| `port`                  | No       | `false`        | No     | `true`, `false`                                                         | The TCP port used to connect to the database server. If not specified, the default port will be used.                                                                                                             |
| `instance`              | No       | *empty string* | No     | *valid instance name*                                                   | The database server instance name. This may be blank.                                                                                                                                                             |
| `db-name`               | Yes      | *empty string* | No     | *valid database name*                                                   | The name of the database used by ILLiad software and checked by plugins from this project.                                                                                                                        |
| `username`              | Yes      | *empty string* | No     | *valid username*                                                        | The username used to connect to the database server. An account with read-only access to the database used by the ILLiad software is sufficient.                                                                  |
| `password`              | Yes      | *empty string* | No     | *valid password*                                                        | The plaintext password used to connect to the database server. An account with read-only access to the database used by the ILLiad software is sufficient.                                                        |
| `encrypt-mode`          | No       | `false`        | No     | `true`, `false`, `disable`                                              | Whether data sent between client and server is encrypted. `true` for yes, `false` for login packet only and `disable` for no encryption.                                                                          |
| `trust-cert`            | No       | `false`        | No     | `true`, `false`                                                         | Whether the certificate should be trusted as-is without validation. WARNING: TLS is susceptible to man-in-the-middle attacks if enabling this option.                                                             |
| `ignore-missing-emails` | No       | `false`        | No     | `true`, `false`                                                         | Whether finding zero email notifications recorded in the database used by ILLiad software should be treated as an `OK` state. Legitimate scenarios include fresh ILLiad installations or recent purge of history. |
| `count-warning`         | No       | `1`            | No     | `1+` (*minimum of 1*)                                                   | The number of pending email notifications when this plugin will consider the service check to be in a `WARNING` state.                                                                                            |
| `count-critical`        | No       | `3`            | No     | `2+` (*minimum 1 greater than `warning`*)                               | The number of pending email notifications when this plugin will consider the service check to be in a `CRITICAL` state.                                                                                           |
| `age-warning`           | No       | `5`            | No     | `1+` (*minimum of 1*)                                                   | The number of minutes an email notification has been in a pending status when this plugin will consider the service check to be in a `WARNING` state.                                                             |
| `age-critical`          | No       | `10`           | No     | `2+` (*minimum 1 greater than `warning`*)                               | The number of minutes an email notification has been in a pending status when this plugin will consider the service check to be in a `CRITICAL` state.                                                            |

### Environment Variables

If used, command-line arguments override the equivalent environment variables
listed below. See the [Command-line Arguments](#command-line-arguments) table
for more information.

| Flag Name               | Environment Variable Name            | Notes | Example (mostly using default values)              |
| ----------------------- | ------------------------------------ | ----- | -------------------------------------------------- |
| `emit-branding`         | `CHECK_ILLIAD_EMAILS_EMIT_BRANDING`  |       | `CHECK_ILLIAD_EMAILS_EMIT_BRANDING="false"`        |
| `log-level`             | `CHECK_ILLIAD_EMAILS_LOG_LEVEL`      |       | `CHECK_ILLIAD_EMAILS_LOG_LEVEL="info"`             |
| `host`                  | `CHECK_ILLIAD_DBSERVER_HOST`         |       | `CHECK_ILLIAD_DBSERVER_HOST="mssql52.example.com"` |
| `port`                  | `CHECK_ILLIAD_DBSERVER_PORT`         |       | `CHECK_ILLIAD_DBSERVER_PORT="1433"`                |
| `instance`              | `CHECK_ILLIAD_DBSERVER_INSTANCE`     |       | `CHECK_ILLIAD_DBSERVER_INSTANCE="mssql07"`         |
| `db-name`               | `CHECK_ILLIAD_DATABASE_NAME`         |       | `CHECK_ILLIAD_DATABASE_NAME="ILLData"`             |
| `username`              | `CHECK_ILLIAD_DBSERVER_USERNAME`     |       | `CHECK_ILLIAD_DBSERVER_USERNAME="dbuser"`          |
| `password`              | `CHECK_ILLIAD_DBSERVER_PASSWORD`     |       | `CHECK_ILLIAD_DBSERVER_PASSWORD="dbPasW0rdZ"`      |
| `encrypt-mode`          | `CHECK_ILLIAD_DBSERVER_ENCRYPT_MODE` |       | `CHECK_ILLIAD_DBSERVER_ENCRYPT_MODE="false"`       |
| `trust-cert`            | `CHECK_ILLIAD_TRUST_CERT`            |       | `CHECK_ILLIAD_TRUST_CERT="false"`                  |
| `ignore-missing-emails` | `CHECK_ILLIAD_IGNORE_MISSING_EMAILS` |       | `CHECK_ILLIAD_IGNORE_MISSING_EMAILS="false"`       |
| `count-warning`         | `CHECK_ILLIAD_COUNT_WARNING`         |       | `CHECK_ILLIAD_COUNT_WARNING="1"`                   |
| `count-critical`        | `CHECK_ILLIAD_COUNT_CRITICAL`        |       | `CHECK_ILLIAD_COUNT_CRITICAL="3"`                  |
| `age-warning`           | `CHECK_ILLIAD_AGE_WARNING`           |       | `CHECK_ILLIAD_AGE_WARNING="5"`                     |
| `age-critical`          | `CHECK_ILLIAD_AGE_CRITICAL`          |       | `CHECK_ILLIAD_AGE_CRITICAL="10"`                   |

## Example

### Default thresholds

```ShellSession
$ ./check_illiad_emails --host mssql52.example.com --port 1433 --instance mssql07 --username dbuser --password 'dbPasW0rdZ' --encrypt-mode false --trust-cert --db-name illdata
OK: 0 email notifications in a pending state found without crossing specified thresholds

**ERRORS**

* None

**THRESHOLDS**

* CRITICAL: [Age: 10m0s, Count: 3]
* WARNING: [Age: 5m0s, Count: 1]

**DETAILED INFO**

* Database Connection
** Host: "mssql52.example.com"
** Port: 1433
** Instance: "mssql07"
** Database: "illdata"
** Encryption Mode: "false"
** Trust Server Cert (disable verification): true
```

## License

```license
MIT License

Copyright (c) 2020 Adam Chalkley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Related projects

- <https://github.com/atc0005/go-nagios>
- <https://github.com/atc0005/check-mail>
- <https://github.com/atc0005/check-cert>
- <https://github.com/atc0005/check-path>

## References

- <https://github.com/denisenkom/go-mssqldb>
- <https://github.com/rs/zerolog>
- <https://github.com/alexflint/go-arg>
- <https://github.com/phayes/permbits>
- <https://github.com/atc0005/go-nagios>

- <https://support.atlas-sys.com/hc/en-us/categories/360000716874-ILLiad>
- <https://nagios-plugins.org/doc/guidelines.html>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/check-illiad>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
