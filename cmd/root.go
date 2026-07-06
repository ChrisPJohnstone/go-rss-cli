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

var (
	verbose bool
	toolDir string
)

func init() {
	defaultDir, err := defaultDir()
	cobra.CheckErr(err)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().StringVarP(&toolDir, toolDirArg, "", defaultDir, "Directory to store config & files")
	rootCmd.AddCommand(feed.FeedCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
