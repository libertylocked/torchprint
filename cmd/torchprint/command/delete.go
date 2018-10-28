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
		// TODO: finish this when API lib is updated

		// userid := viper.GetString("userid")
		// token := viper.GetString("token")
		delAll, _ := cmd.LocalFlags().GetBool("all")
		if !delAll {
			if len(args) == 0 {
				cmd.Println("Please the ID of the job to delete! Use \"torchprint rm [job-id]\"")
				os.Exit(1)
			}
		}
		// jobID := args[0]

		// api := torchprint.NewAPI(userid).SetToken(token)
		// resp, err := api.DeletePrintJob()
		// if err != nil {
		// 	cmd.Println(err)
		// 	os.Exit(1)
		// }
	},
}

func init() {
	deleteCmd.Flags().BoolP("all", "a", false, "Delete all jobs")
}
