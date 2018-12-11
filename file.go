package cliutil

import (
	"fmt"
	"path/filepath"
)

type (
	// ExistFile returns true if file exists otherwise false.
	ExistFile func(string) bool
)

// FindFile finds file from wd to the root directory recursively.
func FindFile(wd, name string, existFile ExistFile) (string, error) {
	for {
		p := filepath.Join(wd, name)
		if existFile(p) {
			return p, nil
		}
		if wd == "/" || wd == "" {
			return "", fmt.Errorf("%s is not found", name)
		}
		wd = filepath.Dir(wd)
	}
}
