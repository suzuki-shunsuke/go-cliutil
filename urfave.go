package cliutil

import (
	"github.com/urfave/cli"
)

// ConvErrToExitError converts an error into urfave/cli's ExitError.
func ConvErrToExitError(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*cli.ExitError); ok {
		return err
	}
	return cli.NewExitError(err.Error(), 1)
}

func WrapAction(f func(c *cli.Context) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		return ConvErrToExitError(f(c))
	}
}
