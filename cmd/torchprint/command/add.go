package command

import (
	"os"

	"github.com/libertylocked/torchprint"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a document to printing queue",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("Please specify a file to print! Use \"torchprint add [filename] [options]\"")
			os.Exit(1)
		}
		filename := args[0]

		userid := viper.GetString("userid")
		token := viper.GetString("token")
		color, _ := cmd.LocalFlags().GetBool("color")
		side, _ := cmd.LocalFlags().GetString("side")
		perSide, _ := cmd.LocalFlags().GetInt("perside")
		copies, _ := cmd.LocalFlags().GetInt("copies")
		pageRange, _ := cmd.LocalFlags().GetString("range")

		api := torchprint.NewAPI(userid).SetToken(token)
		options := torchprint.FinishingOptions{
			Mono: !color,
			Duplex: func() bool {
				if side == "single" {
					return false
				}
				return true
			}(),
			PagesPerSide:    perSide,
			Copies:          copies,
			DefaultPageSize: "Letter",
			PageRange:       pageRange,
		}
		resp, err := api.AddPrintJob(filename, options)
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		cmd.Println("success:", resp.Location, resp.PrintState, resp.Name)
	},
}

func init() {
	addCmd.Flags().Bool("color", false, "Set color or monochrome mode")
	addCmd.Flags().String("side", "double", "Set single or double sided (\"single\" or \"double\")")
	addCmd.Flags().Int("perside", 1, "Pages per side")
	addCmd.Flags().Int("copies", 1, "Number of copies to print")
	addCmd.Flags().String("range", "", "Page range (e.g. \"1-2\", \"3\")")
}
