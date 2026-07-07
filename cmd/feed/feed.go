package feed

import "github.com/spf13/cobra"

var FeedCmd = &cobra.Command{
	Use:   "feed [command]",
	Short: "Feed management commands",
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	FeedCmd.AddCommand(listFeedCmd, addFeedCmd, rmFeedCmd)
}
