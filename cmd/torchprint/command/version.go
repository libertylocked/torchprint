package command

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of torchprint",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("torchprint v0.1 -- HEAD")
	},
}
