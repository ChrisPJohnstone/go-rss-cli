package main

import (
	"fmt"
	"log"

	"github.com/ChrisPJohnstone/go-rss-client"
)

func main() {
	opts, err := parseOpts()
	if err != nil {
		log.Fatal(err)
	}
	feed, err := rssclient.FetchFeed(opts.url)
	if err != nil {
		log.Fatal(err)
	}
	titles := make([]string, len(feed.Items))
	for i, item := range feed.Items {
		titles[i] = item.Title
	}
	fmt.Println(len(titles))
}
