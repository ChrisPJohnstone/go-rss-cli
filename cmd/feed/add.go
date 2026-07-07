package feed

import (
	"fmt"

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
	fmt.Println(internal.Cfg)
	return fmt.Errorf("not implemented")
}
