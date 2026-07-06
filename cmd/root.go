package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/ChrisPJohnstone/go-rss-cli/cmd/feed"
	"github.com/ChrisPJohnstone/go-rss-cli/internal"
)

var rootCmd = &cobra.Command{
	Use:               "rss",
	Short:             "Fast & Minimal Feed Reader",
	PersistentPreRunE: internal.LoadConfig,
}

func init() {
	defaultDir, err := internal.DefaultDir()
	cobra.CheckErr(err)
	rootCmd.PersistentFlags().BoolVarP(&internal.Verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().StringVarP(&internal.ToolDir, internal.ToolDirArg, "", defaultDir, "Directory to store config & files")
	rootCmd.AddCommand(feed.FeedCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
