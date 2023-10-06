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

## [v0.2.5] - 2023-10-06

### Changed

#### Dependency Updates

- (GH-250) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-240) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.5 to go-ci-oldstable-build-v0.13.6 in /dependabot/docker/builds
- (GH-242) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.6 to go-ci-oldstable-build-v0.13.7 in /dependabot/docker/builds
- (GH-252) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.7 to go-ci-oldstable-build-v0.13.8 in /dependabot/docker/builds
- (GH-258) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.8 to go-ci-oldstable-build-v0.13.9 in /dependabot/docker/builds
- (GH-247) ghaw: bump actions/checkout from 3 to 4
- (GH-248) go.mod: bump golang.org/x/crypto from 0.12.0 to 0.13.0
- (GH-244) go.mod: bump golang.org/x/sys from 0.11.0 to 0.12.0

## [v0.2.4] - 2023-08-21

### Added

- (GH-207) Add initial automated release notes config
- (GH-209) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.3` to `go-ci-oldstable-build-v0.13.5`
  - `rs/zerolog`
    - `v1.29.1` to `v1.30.0`
  - `golang.org/x/crypto`
    - `v0.11.0` to `v0.12.0`
  - `golang.org/x/sys`
    - `v0.10.0` to `v0.11.0`
- (GH-211) Update Dependabot config to monitor both branches
- (GH-233) Update project to Go 1.20 series

## [v0.2.3] - 2023-07-13

### Overview

- **NOTE**: Change exit state for connection failure
- RPM package improvements
- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.10` to `1.19.11`
  - `atc0005/go-nagios`
    - `v0.15.0` to `v0.16.0`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.6` to `go-ci-oldstable-build-v0.11.3`
  - `golang.org/x/crypto`
    - `v0.10.0` to `v0.11.0`
  - `golang.org/x/sys`
    - `v0.9.0` to `v0.10.0`
- (GH-202) Update RPM postinstall scripts to use restorecon

### Fixed

- (GH-199) Use CRITICAL state for connection failures

## [v0.2.2] - 2023-06-14

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- Dependency updates
- built using Go 1.19.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.9` to `1.19.10`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.5` to `go-ci-oldstable-build-v0.10.6`
  - `atc0005/go-nagios`
    - `v0.14.0` to `v0.15.0`
  - `mattn/go-isatty`
    - `v0.0.18` to `v0.0.19`
  - `golang.org/x/crypto`
    - `v0.9.0` to `v0.10.0`
  - `golang.org/x/sys`
    - `v0.8.0` to `v0.9.0`
- (GH-191) Update vuln analysis GHAW to remove on.push hook

### Fixed

- (GH-188) Disable depguard linter
- (GH-194) Restore local CodeQL workflow

## [v0.2.1] - 2023-05-12

### Overview

- Dependency updates
- built using Go 1.19.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.8` to `1.19.9`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.4` to `go-ci-oldstable-build-v0.10.5`
  - `rs/zerolog`
    - `v1.29.0` to `v1.29.1`
  - `golang.org/x/sys`
    - `v0.7.0` to `v0.8.0`
  - `golang.org/x/crypto`
    - `v0.7.0` to `v0.9.0`

## [v0.2.0] - 2023-04-06

### Overview

- Add support for generating DEB, RPM packages
- Build improvements
- Generated binary changes
  - filename patterns
  - compression (~ 66% smaller)
  - executable metadata
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-161) Generate RPM/DEB packages using nFPM
- (GH-162) Add version details to Windows executables

### Changed

- (GH-160) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-163) Makefile: Compress binaries & use fixed filenames
- (GH-164) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-165) Build dev/stable releases using go-ci Docker image

## [v0.1.16] - 2023-04-06

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-148) Add Go Module Validation, Dependency Updates jobs

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.8`
  - `atc0005/go-nagios`
    - `v0.10.2` to `v0.14.0`
  - `rs/zerolog`
    - `v1.28.0` to `v1.29.0`
  - `github.com/mattn/go-isatty`
    - `v0.0.16` to `v0.0.18`
  - `golang.org/x/crypto`
    - `v0.3.0` to `v0.7.0`
  - `golang.org/x/sys`
    - `v0.3.0` to `v0.7.0`
- CI
  - (GH-155) Drop `Push Validation` workflow
  - (GH-156) Rework workflow scheduling
  - (GH-158) Remove `Push Validation` workflow status badge
- Misc
  - (GH-143) Update nagios library usage, add time perfdata

### Fixed

- (GH-169) Update vuln analysis GHAW to use on.push hook
- (GH-173) Use UNKNOWN state for invalid command-line args
- (GH-174) Use UNKNOWN state for email evaluation failures

## [v0.1.15] - 2022-12-09

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.1` to `1.19.4`
  - `atc0005/go-nagios`
    - `v0.10.0` to `v0.10.2`
  - `github.com/denisenkom/go-mssqldb`
    - `v0.12.2` to `v0.12.3`
  - `github.com/mattn/go-colorable`
    - `v0.1.12` to `v0.1.13`
  - `github.com/mattn/go-isatty`
    - `v0.0.14` to `v0.0.16`
  - `github.com/golang-sql/civil`
    - `v0.0.0-20190719163853-cb61b32ac6fe` to
      `v0.0.0-20220223132316-b832511892a9`
  - `github.com/alexflint/go-scalar`
    - `v1.1.0` to `v1.2.0`
  - `golang.org/x/crypto`
    - `v0.0.0-20220622213112-05595931fe9d` to `v0.3.0`
  - `golang.org/x/sys`
    - `v0.0.0-20210927094055-39ccf1dd6fa6` to `v0.3.0`
- (GH-129) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-136) Fix Makefile Go module base path detection

## [v0.1.14] - 2022-09-21

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.1
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.13` to `1.19.1`
  - `atc0005/go-nagios`
    - `v0.9.1` to `v0.10.0`
  - `rs/zerolog`
    - `v1.27.0` to `v1.28.0`
  - `github/codeql-action`
    - `v2.1.21` to `v2.1.24`
- (GH-119) Update project to Go 1.19
- (GH-120) Update Makefile and GitHub Actions Workflows

### Fixed

- (GH-118) Add missing cmd doc file

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

[Unreleased]: https://github.com/atc0005/check-illiad/compare/v0.2.5...HEAD
[v0.2.5]: https://github.com/atc0005/check-illiad/releases/tag/v0.2.5
[v0.2.4]: https://github.com/atc0005/check-illiad/releases/tag/v0.2.4
[v0.2.3]: https://github.com/atc0005/check-illiad/releases/tag/v0.2.3
[v0.2.2]: https://github.com/atc0005/check-illiad/releases/tag/v0.2.2
[v0.2.1]: https://github.com/atc0005/check-illiad/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/check-illiad/releases/tag/v0.2.0
[v0.1.16]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.16
[v0.1.15]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.15
[v0.1.14]: https://github.com/atc0005/check-illiad/releases/tag/v0.1.14
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
