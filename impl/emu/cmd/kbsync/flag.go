package main

import (
	"flag"
)

var (
	flagOutput string
	flagDSID   uint64
)

func init() {
	flag.CommandLine.StringVar(&flagOutput, "o", "-", "The output file path")
	flag.CommandLine.Uint64Var(&flagDSID, "dsid", 0, "The directory services identifier")
	flag.Parse()
}
