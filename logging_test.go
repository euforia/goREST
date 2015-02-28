package goREST

import (
	"os"
	"testing"
)

var testLogger = NewLogger(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stderr)

func Test_NewStdLogger(t *testing.T) {
	NewStdLogger()
}

func Test_NewLogger(t *testing.T) {
	testLogger.SetLogLevel("error")
	testLogger.SetLogLevel("warning")
	testLogger.SetLogLevel("info")
	testLogger.SetLogLevel("debug")
	testLogger.SetLogLevel("trace")
}

func Test_NewStdLogger_Error(t *testing.T) {
	if err := testLogger.SetLogLevel("invalid"); err == nil {
		t.Fatalf("Error check failed!")
	}
}
