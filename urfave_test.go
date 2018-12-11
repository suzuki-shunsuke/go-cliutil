package cliutil

import (
	"testing"
)

func TestConvErrToExitError(t *testing.T) {
	if err := ConvErrToExitError(nil); err != nil {
		t.Fatal(err)
	}
}
