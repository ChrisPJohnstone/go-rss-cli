package feed

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/go-rss-cli/internal"
)

var addFeedCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Follow a feed",
	// TODO: Support multi-arg
	Args: cobra.MinimumNArgs(1),
	RunE: addFeed,
}

func addFeed(cmd *cobra.Command, args []string) error {
	url := args[0]
	if internal.Verbose {
		log.Printf("Adding url %s to feeds", url)
	}
	if hasFeed(url) {
		return fmt.Errorf("feed %q already exists", url)
	}
	internal.Cfg.Feeds = append(internal.Cfg.Feeds, url)
	err := internal.WriteConfig()
	return err
}
