package feed

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/go-rss-cli/internal"
)

var listFeedCmd = &cobra.Command{
	Use:   "ls",
	Short: "List followed feeds",
	RunE:  listFeeds,
}

func listFeeds(cmd *cobra.Command, args []string) error {
	for _, f := range internal.Cfg.Feeds {
		fmt.Println(f)
	}
	return nil
}
