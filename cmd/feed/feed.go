package feed

import (
	"fmt"

	"github.com/spf13/cobra"
)

var FeedCmd = &cobra.Command{
	Use:   "feed [command]",
	Short: "Feed management commands",
	Args:  cobra.MaximumNArgs(1),
}

var addFeedCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Add a new feed",
	Args:  cobra.MinimumNArgs(1),
	RunE:  addFeed,
}

var rmFeedCmd = &cobra.Command{
	Use:   "rm [url]",
	Short: "Remove a feed",
	Args:  cobra.MinimumNArgs(1),
	RunE:  rmFeed,
}

func init() {
	FeedCmd.AddCommand(addFeedCmd, rmFeedCmd)
}

func addFeed(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("not implemented")
}

func rmFeed(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("not implemented")
}
