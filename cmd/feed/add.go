package feed

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addFeedCmd = &cobra.Command{
	Use:   "add [url]",
	Short: "Add a new feed",
	Args:  cobra.MinimumNArgs(1),
	RunE:  addFeed,
}

func addFeed(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("not implemented")
}
