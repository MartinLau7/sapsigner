package main

import (
	"os"
)

func openOr(name string, f *os.File) (*os.File, error) {
	if name == "-" {
		return f, nil
	}

	return os.Open(name)
}
