package feed

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/go-rss-cli/internal"
)

var rmFeedCmd = &cobra.Command{
	Use:   "rm [url]",
	Short: "Unfollow a feed",
	// TODO: Support multi-arg
	Args: cobra.MinimumNArgs(1),
	RunE: rmFeed,
}

func rmFeed(cmd *cobra.Command, args []string) error {
	url := args[0]
	if internal.Verbose {
		log.Printf("Removing url %s from feeds", url)
	}
	if !hasFeed(url) {
		return fmt.Errorf("feed %q not found", url)
	}
	var updated []string
	for _, f := range internal.Cfg.Feeds {
		if f != url {
			updated = append(updated, f)
		}
	}
	internal.Cfg.Feeds = updated
	return internal.WriteConfig()
}
