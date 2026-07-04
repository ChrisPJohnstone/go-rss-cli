package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Config struct {
	Verbose bool
	URL     string
}

func parseConfig() (Config, error) {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.Parse()
	config := Config{Verbose: verbose}
	if config.Verbose {
		log.Printf("Verbose logging enabled")
	}
	stat, _ := os.Stdin.Stat()
	hasStdin := (stat.Mode() & os.ModeCharDevice) == 0
	args := flag.Args()
	if hasStdin && len(args) > 0 {
		return Config{}, fmt.Errorf("accepts input from stdin OR positional arg, not both")
	}
	var input []string
	if hasStdin {
		if config.Verbose {
			log.Printf("Reading input from stdin")
		}
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return Config{}, err
		}
		input = strings.Fields(string(bytes))
	} else {
		if config.Verbose {
			log.Printf("Reading input from positional args")
		}
		input = args
	}
	// TODO: Depending how implementation shakes out could accept multiple inputs
	if len(input) > 1 {
		return Config{}, fmt.Errorf("expected one positional arg")
	}
	config.URL = input[0]
	return config, nil
}

func run(config Config) error {
	if config.Verbose {
		log.Printf("input: %v", config)
	}
	return nil
}

func main() {
	config, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := run(config); err != nil {
		log.Fatal(err)
	}
}
