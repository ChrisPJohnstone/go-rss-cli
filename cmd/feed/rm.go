package feed

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rmFeedCmd = &cobra.Command{
	Use:   "rm [url]",
	Short: "Remove a feed",
	Args:  cobra.MinimumNArgs(1),
	RunE:  rmFeed,
}

func rmFeed(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("not implemented")
}
