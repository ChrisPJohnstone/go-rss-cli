package cmd

import (
	"fmt"

	"github.com/ChrisPJohnstone/go-rss-client"
	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/go-rss-cli/internal"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull articles from followed feeds",
	RunE:  pullAll,
}

func pullFeed(url string) error {
	feed, err := rssclient.FetchFeed(url)
	if err != nil {
		return err
	}
	for _, item := range feed.Items {
		// TODO: Write somewhere
		fmt.Println(item.Title)
	}
	return nil
}

func pullAll(cmd *cobra.Command, args []string) error {
	for _, f := range internal.Cfg.Feeds {
		// TODO: Make async
		err := pullFeed(f)
		if err != nil {
			return err
		}
	}
	return nil
}
