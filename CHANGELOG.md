# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/check-illiad/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.1.3] - 2021-07-16

### Overview

- Add test
- Dependency updates
- built using Go 1.16.6
  - Statically linked
  - Linux (x86, x64)

### Added

- Add test
  - basic custom flags parsing using example from README

- Add "canary" Dockerfile to track stable Go releases, serve as a reminder to
  generate fresh binaries

### Changed

- dependencies
  - `Go`
    - `1.16.3` to `1.16.6`
    - canary file updated from `1.16.5` to `1.16.6`
  - `atc0005/go-nagios`
    - `v0.6.0` to `v0.6.1`
  - `alexflint/go-arg`
    - `v1.3.0` to `v1.4.2`
  - `pelletier/go-toml`
    - `v1.9.2` to `v1.9.3`
  - `rs/zerolog`
    - `v1.22.0` to `v1.23.0`
  - `actions/setup-node`
    - updated from `v2.1.5` to `v2.2.0`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

## [v0.1.2] - 2021-04-15

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.16.3

### Changed

- dependencies
  - built using `Go 1.16.3`
    - Windows (x86, x64)
    - Linux (x86, x64)
  - `denisenkom/go-mssqldb`
    - `v0.0.0-20201104001602-427686ac8ec1` to `v0.10.0`
  - `rs/zerolog`
    - `v1.20.0` to `v1.21.0`
  - `actions/setup-node`
    - `v2.1.4` to `v2.1.5`

### Fixed

- Fix description of possible values for host flag
- Fix invalid db and username debug output
- Add missing newline between bullet points

## [v0.1.1] - 2021-02-07

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.15.8

### Changed

- dependencies
  - `atc0005/go-nagios`
    - `v0.5.1` to `v0.6.0`
  - `actions/setup-node`
    - `v2.1.2` to `v2.1.4`
  - Built using `Go 1.15.8`
    - Windows (x86, x64)
    - Linux (x86, x64)
- Remove temporary workaround for swallowed panics
- Replace godoc.org badge with pkg.go.dev badge

### Fixed

- `sql: Scan error on column index 3, name "Note": converting NULL to string
  is unsupported`

## [v0.1.0] - 2020-11-05

### Added

Initial release!

This release provides an early version of a Nagios plugin used to check the
status of email notifications for an ILLiad instance. In particular, this
plugin looks for notifications in a `Pending` status that have remained that
way at a specific number and length of time. Default values may be usable
as-is, otherwise are easily overridden with command-line flags. See the README
file for additional details.

- Statically linked binary release
  - Built using Go 1.15.3
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

[Unreleased]: https://github.com/atc0005/check-illiad/compare/v0.1.3...HEAD
[v0.1.3]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.0
