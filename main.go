package main

import (
	"fmt"
	"log"

	"github.com/ChrisPJohnstone/go-rss-client"
)

func main() {
	var url string = "https://feeds.bbci.co.uk/news/rss.xml"
	feed, err := rssclient.FetchFeed(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)
}
