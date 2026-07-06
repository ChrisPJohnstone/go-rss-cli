package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/go-rss-cli/cmd/feed"
)

var rootCmd = &cobra.Command{
	Use:   "rss",
	Short: "Fast & Minimal Feed Reader",
}

func init() {
	rootCmd.AddCommand(feed.FeedCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
