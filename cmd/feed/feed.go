package feed

import "github.com/spf13/cobra"

var FeedCmd = &cobra.Command{
	Use:   "feed [command]",
	Short: "Feed management commands",
	Args:  cobra.MaximumNArgs(1),
}

var addFeedCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Add a new feed",
	Args:  cobra.MinimumNArgs(1),
	Run:   addFeed,
}

var rmFeedCmd = &cobra.Command{
	Use:   "rm [url]",
	Short: "Remove a feed",
	Args:  cobra.MinimumNArgs(1),
	Run:   rmFeed,
}

func init() {
	FeedCmd.AddCommand(addFeedCmd, rmFeedCmd)
}

func addFeed(cmd *cobra.Command, args []string) {
}

func rmFeed(cmd *cobra.Command, args []string) {
}
