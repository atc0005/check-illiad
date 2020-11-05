// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/check-illiad
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

/*

Go-based tooling to check/verify an ILLiad server instance.

PROJECT HOME

See our GitHub repo (https://github.com/atc0005/check-illiad) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

PURPOSE

This repo is intended to provide Nagios plugins that may be used to monitor an
ILLiad instance. As of this writing, a plugin named check_illiad_emails is
available that may be used to monitor email notifications.

FEATURES

The check_illiad_emails plugin supports monitoring Pending email
notifications:

• greater number queued than specified
• that have remained "in the queue" for longer than a specified amount of time

Default values are provided, but are easily overridden with custom values in
order to match usage patterns for each instance.

USAGE

See our main README for supported settings and examples.

*/
package main
