package main

import (
	"flag"
)

var (
	flagInput  string
	flagOutput string
	flagPrimed bool
)

func init() {
	flag.CommandLine.StringVar(&flagInput, "i", "-", "The input file path")
	flag.CommandLine.StringVar(&flagOutput, "o", "-", "The output file path")
	flag.CommandLine.BoolVar(&flagPrimed, "p", false, "Use primed signing session")
	flag.Parse()
}
