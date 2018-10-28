package command

import (
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete a job from printjob queue",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println("Please the ID of the job to delete! Use \"torchprint rm [job-id]\"")
			os.Exit(1)
		}
		// TODO: finish this when API lib is updated
		// jobID := args[0]

		// userid := viper.GetString("userid")
		// token := viper.GetString("token")

		// api := torchprint.NewAPI(userid).SetToken(token)
		// resp, err := api.DeletePrintJob()
		// if err != nil {
		// 	cmd.Println(err)
		// 	os.Exit(1)
		// }
	},
}
