// config.go
//
// This file implements the configuration part for when you need the API
// key to modify things in the Atlas configuration and manage measurements.

package main

import (
	"path/filepath"
	"strings"
	"os"
)

var (
	// Default location is now $HOME/.config/<tag>/ on UNIX
	basedir = filepath.Join(os.Getenv("HOME"), ".config", MyName)

	// That one is common to all connected utilities
	dbrcFile = filepath.Join(os.Getenv("HOME"), ".dbrc")
)

// Check the parameter for either tag or filename
func checkName(file string) (str string) {
	// Full path, MUST have .toml
	if bfile := []byte(file); bfile[0] == '/' {
		if !strings.HasSuffix(file, ".toml") {
			str = ""
		} else {
			str = file
		}
		return
	}

	// If ending with .toml, take it literally
	if strings.HasSuffix(file, ".toml") {
		str = file
		return
	}

	// Check for tag
	if !strings.HasSuffix(file, ".toml") {
		str = filepath.Join(basedir, file, "config.toml")
	}
	return
}
