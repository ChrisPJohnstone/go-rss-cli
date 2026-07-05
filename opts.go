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
	verbose bool
	dir     string
	url     string
}

func defaultDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".rss-cli"), nil
}

func parseOpts() (opts, error) {
	var verbose bool
	var dir string
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.StringVar(&dir, "dir", "", "directory to store config & files")
	flag.Parse()
	o := opts{verbose: verbose}
	if o.verbose {
		log.Printf("verbose logging enabled")
	}
	if dir == "" {
		if o.verbose {
			log.Printf("getting default directory")
		}
		defaultDir, err := defaultDir()
		if err != nil {
			return opts{}, err
		}
		dir = defaultDir
	}
	o.dir = dir
	if o.verbose {
		log.Printf("directory set to %s", o.dir)
	}
	stat, _ := os.Stdin.Stat()
	hasStdin := (stat.Mode() & os.ModeCharDevice) == 0
	args := flag.Args()
	if hasStdin && len(args) > 0 {
		return opts{}, fmt.Errorf("accepts input from stdin OR positional arg, not both")
	}
	var input []string
	if hasStdin {
		if o.verbose {
			log.Printf("reading input from stdin")
		}
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return opts{}, err
		}
		input = strings.Fields(string(bytes))
	} else {
		if o.verbose {
			log.Printf("reading input from positional args")
		}
		input = args
	}
	// TODO: Depending how implementation shakes out could accept multiple inputs
	if len(input) > 1 {
		return opts{}, fmt.Errorf("expected one positional arg")
	}
	o.url = input[0]
	return o, nil
}
