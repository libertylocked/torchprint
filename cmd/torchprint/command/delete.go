package command

import (
	"os"

	"github.com/libertylocked/torchprint"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete a job from printjob queue",
	Run: func(cmd *cobra.Command, args []string) {
		userid := viper.GetString("userid")
		token := viper.GetString("token")
		delAll, _ := cmd.LocalFlags().GetBool("all")

		api := torchprint.NewAPI(userid).SetToken(token)
		deleteLocations := []torchprint.PrintJobDeleteLocation{}

		if delAll {
			jobs, err := api.GetPrintJobs()
			if err != nil {
				cmd.Println("error:", "fail to delete all printjobs - cannot get printjobs:", err)
				os.Exit(1)
			}
			if jobs != nil {
				for _, job := range jobs.Items {
					deleteLocations = append(deleteLocations, torchprint.PrintJobDeleteLocation{
						Location: job.Location,
					})
				}
			}
		} else {
			if len(args) == 0 {
				cmd.Println("Please the ID of the job to delete! Use \"torchprint rm [job-id]\"")
				os.Exit(1)
			}
			for _, jobID := range args {
				deleteLocations = append(deleteLocations, torchprint.PrintJobDeleteLocation{
					Location: "/printjobs/" + jobID,
				})
			}
		}

		resp, err := api.DeletePrintJobs(deleteLocations)
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}
		// pretty print response
		for _, deleteStatus := range resp {
			cmd.Println(deleteStatus.Location, deleteStatus.Status)
		}
	},
}

func init() {
	deleteCmd.Flags().BoolP("all", "a", false, "Delete all jobs")
}
