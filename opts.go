package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type opts struct {
	verbose bool
	url     string
}

func parseOpts() (opts, error) {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.Parse()
	config := opts{verbose: verbose}
	if config.verbose {
		log.Printf("verbose logging enabled")
	}
	stat, _ := os.Stdin.Stat()
	hasStdin := (stat.Mode() & os.ModeCharDevice) == 0
	args := flag.Args()
	if hasStdin && len(args) > 0 {
		return opts{}, fmt.Errorf("accepts input from stdin OR positional arg, not both")
	}
	var input []string
	if hasStdin {
		if config.verbose {
			log.Printf("reading input from stdin")
		}
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return opts{}, err
		}
		input = strings.Fields(string(bytes))
	} else {
		if config.verbose {
			log.Printf("reading input from positional args")
		}
		input = args
	}
	// TODO: Depending how implementation shakes out could accept multiple inputs
	if len(input) > 1 {
		return opts{}, fmt.Errorf("expected one positional arg")
	}
	config.url = input[0]
	return config, nil
}
