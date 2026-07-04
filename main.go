package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	// "github.com/ChrisPJohnstone/go-rss-client"
)

func getInput() ([]string, error) {
	stat, _ := os.Stdin.Stat()
	hasStdin := (stat.Mode() & os.ModeCharDevice) == 0
	args := flag.Args()
	if hasStdin && len(args) > 0 {
		return nil, fmt.Errorf("accepts input from stdin OR positional arg, not both")
	}
	if hasStdin {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return strings.Fields(string(bytes)), nil
	}
	return args, nil
}

func main() {
	flag.Parse()
	input, err := getInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
	// TODO: parse input

	// var url string = "https://feeds.bbci.co.uk/news/rss.xml"
	// feed, err := rssclient.FetchFeed(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(feed)
}
