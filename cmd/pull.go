package cmd

import (
	"github.com/ChrisPJohnstone/go-rss-client/v2"
	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/go-rss-cli/internal"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull articles from followed feeds",
	RunE:  pullAll,
}

func pullFeed(url string) error {
	feed, _, err := rssclient.FetchFeed(url, nil)
	// TODO: Use ETag & LastModified to pass when no changes
	if err != nil {
		return err
	}
	for _, item := range feed.Items {
		internal.InsertArticle(url, item)
	}
	return nil
}

func pullAll(cmd *cobra.Command, args []string) error {
	err := internal.InitDB()
	if err != nil {
		return err
	}
	for _, f := range internal.Cfg.Feeds {
		// TODO: Make async
		err := pullFeed(f)
		if err != nil {
			return err
		}
	}
	return nil
}
