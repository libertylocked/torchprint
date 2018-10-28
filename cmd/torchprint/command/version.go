package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of torchprint",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("torchprint v0.1 -- HEAD")
	},
}
