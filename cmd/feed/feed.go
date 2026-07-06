package feed

import "github.com/spf13/cobra"

var FeedCmd = &cobra.Command{
	Use:               "feed [command]",
	Short:             "Feed management commands",
	Args:              cobra.MaximumNArgs(1),
	PersistentPreRunE: loadSharedFlags,
}

var (
	verbose bool
	toolDir string
)

func init() {
	FeedCmd.AddCommand(addFeedCmd, rmFeedCmd)
}

func loadSharedFlags(cmd *cobra.Command, args []string) error {
	var err error
	toolDir, err = cmd.Flags().GetString("toolDir")
	if err != nil {
		return err
	}
	verbose, err = cmd.Flags().GetBool("verbose")
	if err != nil {
		return err
	}
	return nil
}
