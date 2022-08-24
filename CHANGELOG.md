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

## [v0.1.13] - 2022-08-24

### Overview

- Dependency updates
- Bug fixes
- built using Go 1.17.13
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.12` to `1.17.13`

### Fixed

- (GH-115) Apply Go 1.19 specific doc comments linting fixes

## [v0.1.12] - 2022-07-21

### Overview

- Dependency updates
- Bugfixes
- built using Go 1.17.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.10` to `1.17.12`
  - `denisenkom/go-mssqldb`
    - `v0.12.0` to `v0.12.2`
  - `atc0005/go-nagios`
    - `v0.8.2` to `v0.9.1`
  - `rs/zerolog`
    - `v1.26.1` to `v1.27.0`

### Fixed

- (GH-110) Update lintinstall Makefile recipe
- (GH-111) Fix various atc0005/go-nagios usage linting errors

## [v0.1.11] - 2022-05-13

### Overview

- Dependency updates
- built using Go 1.17.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.9` to `1.17.10`

## [v0.1.10] - 2022-05-05

### Overview

- Dependency updates
- built using Go 1.17.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.17.9`

## [v0.1.9] - 2022-03-02

### Overview

- Dependency updates
- CI / linting improvements
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
  - `denisenkom/go-mssqldb`
    - `v0.11.0` to `v0.12.0`
  - `alexflint/go-arg`
    - `v1.4.2` to `v1.4.3`
  - `actions/checkout`
    - `v2.4.0` to `v3`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-89) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-90) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

### Fixed

- (GH-92) gosec, revive linting errors surfaced by GHAWs refresh

## [v0.1.8] - 2022-01-21

### Overview

- Dependency updates
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.12` to `1.17.6`
    - (GH-84) Update go.mod file, canary Dockerfile to reflect current
      dependencies
  - `atc0005/go-nagios`
    - `v0.8.1` to `v0.8.2`

## [v0.1.7] - 2021-12-29

### Overview

- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
  - `rs/zerolog`
    - `v1.26.0` to `v1.26.1`
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.1`

## [v0.1.6] - 2021-11-09

### Overview

- Dependency updates
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.8` to `1.16.10`
  - `atc0005/go-nagios`
    - `v0.7.0` to `v0.8.1`
  - `rs/zerolog`
    - `v1.25.0` to `v1.26.0`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

## [v0.1.5] - 2021-09-25

### Overview

- Dependency updates
- built using Go 1.16.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.8`
  - `denisenkom/go-mssqldb`
    - `v0.10.0` to `v0.11.0`
  - `atc0005/go-nagios`
    - `v0.6.1` to `v0.7.0`
  - `rs/zerolog`
    - `v1.23.0` to `v1.25.0`

## [v0.1.4] - 2021-08-08

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
  - `actions/setup-node`
    - updated from `v2.2.0` to `v2.4.0`

### Fixed

- README
  - Tweak wording of queue count note

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

[Unreleased]: https://github.com/atc0005/check-illiad/compare/v0.1.13...HEAD
[v0.1.13]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.13
[v0.1.12]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.12
[v0.1.11]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.11
[v0.1.10]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.10
[v0.1.9]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.9
[v0.1.8]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.8
[v0.1.7]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.7
[v0.1.6]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.6
[v0.1.5]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.5
[v0.1.4]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.4
[v0.1.3]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.0
