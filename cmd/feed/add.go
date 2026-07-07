package feed

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/go-rss-cli/internal"
)

var addFeedCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Add a new feed",
	Args:  cobra.MinimumNArgs(1),
	RunE:  addFeed,
}

func addFeed(cmd *cobra.Command, args []string) error {
	url := args[0]
	if internal.Verbose {
		log.Printf("Adding url %s to feeds", url)
	}
	// TODO: Check for duplicates
	internal.Cfg.Feeds = append(internal.Cfg.Feeds, url)
	err := internal.WriteConfig()
	return err
}
