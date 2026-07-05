package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type opts struct {
	verbose    bool
	configPath string
	url        string
}

func parseOpts() (opts, error) {
	var verbose bool
	var configPath string
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.StringVar(&configPath, "config-path", "", "path to config file")
	flag.Parse()
	config := opts{verbose: verbose}
	if config.verbose {
		log.Printf("verbose logging enabled")
	}
	if configPath == "" {
		if config.verbose {
			log.Printf("getting default config path")
		}
		home, err := os.UserHomeDir()
		if err != nil {
			return opts{}, fmt.Errorf("cannot determine home directory: %w", err)
		}
		configPath = filepath.Join(home, ".rss-cli", "config")
	}
	config.configPath = configPath
	if config.verbose {
		log.Printf("config path set to %s", config.configPath)
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
