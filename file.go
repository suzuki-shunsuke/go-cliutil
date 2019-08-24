package cliutil

import (
	"errors"
	"path/filepath"
)

type (
	// ExistFile returns true if file exists otherwise false.
	ExistFile func(string) bool
)

// FindFile finds file from wd to the root directory recursively.
func FindFile(wd string, existFile ExistFile, names ...string) (string, error) {
	for {
		for _, name := range names {
			p := filepath.Join(wd, name)
			if existFile(p) {
				return p, nil
			}
		}
		if wd == "/" || wd == "" {
			return "", errors.New("file is not found")
		}
		wd = filepath.Dir(wd)
	}
}
