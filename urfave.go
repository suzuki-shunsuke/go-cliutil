package cliutil

import (
	"github.com/urfave/cli"
)

// ConvErrToExitError converts an error into urfave/cli's ExitError.
func ConvErrToExitError(err error) error {
	if err == nil {
		return nil
	}
	return cli.NewExitError(err.Error(), 1)
}
