// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/check-illiad
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/atc0005/check-illiad/internal/config"
	"github.com/atc0005/go-nagios"
)

// Setup basic tests to ensure that alexflint/go-arg package recognizes flags
// (see atc0005/query-meta#15 and alexflint/go-arg#157 for backstory) and that
// config file loading works as expected.

func TestFlagsParsing(t *testing.T) {

	// https://stackoverflow.com/questions/33723300/how-to-test-the-passing-of-arguments-in-golang

	// Save old command-line arguments so that we can restore them later
	oldArgs := os.Args

	// Defer restoring original command-line arguments
	defer func() { os.Args = oldArgs }()

	// Note to self: Don't add/escape double-quotes here. The shell strips
	// them away and the application never sees them.
	os.Args = []string{
		"/usr/lib/nagios/plugins/check_illiad_emails",
		"--host", "mssql52.example.com",
		"--port", "1433",
		"--instance", "mssql07",
		"--username", "dbuser",
		"--password", "dbPasW0rdZ",
		"--encrypt-mode", "false",
		"--trust-cert",
		"--db-name", "illdata",
	}

	_, err := config.New()
	if err != nil {
		t.Errorf("Error encountered when instantiating configuration: %s", err)
	} else {
		t.Log("No errors encountered when instantiating configuration")
	}

}

// TestEmptyClientPerfDataAndConstructedPluginProducesDefaultTimeMetric
// asserts that omitted performance data from client code produces a default
// time metric when using the Plugin constructor.
func TestEmptyClientPerfDataAndConstructedPluginProducesDefaultTimeMetric(t *testing.T) {
	t.Parallel()

	// Setup Plugin type the same way that client code using the
	// constructor would.
	plugin := nagios.NewPlugin()

	// Performance Data metrics are not emitted if we do not supply a
	// ServiceOutput value.
	plugin.ServiceOutput = "TacoTuesday"

	var outputBuffer strings.Builder

	plugin.SetOutputTarget(&outputBuffer)

	// os.Exit calls break tests
	plugin.SkipOSExit()

	// Process exit state, emit output to our output buffer.
	plugin.ReturnCheckResults()

	want := fmt.Sprintf(
		"%s | %s",
		plugin.ServiceOutput,
		"'time'=",
	)

	got := outputBuffer.String()

	if !strings.Contains(got, want) {
		t.Errorf("ERROR: Plugin output does not contain the expected time metric")
		t.Errorf("\nwant %q\ngot %q", want, got)
	} else {
		t.Logf("OK: Emitted performance data contains the expected time metric.")
	}
}
