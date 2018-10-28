package command

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"text/tabwriter"

	"github.com/libertylocked/torchprint"
	"github.com/libertylocked/urlpattern"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List jobs in printjob queue",
	Run: func(cmd *cobra.Command, args []string) {
		userid := viper.GetString("userid")
		token := viper.GetString("token")
		api := torchprint.NewAPI(userid).SetToken(token)
		resp, err := api.GetPrintJobs()
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}

		// pretty print jobs
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 20, 4, 3, ' ', 0)
		fmt.Fprintln(w, "JOB ID\tNAME\tSUBMISSION TIME\tSTATE")
		pattern := urlpattern.NewPattern().Path("/printjobs/{jobid}")
		for _, job := range resp.Items {
			u, _ := url.Parse(job.Location)
			matches, _ := pattern.Match(u)
			fmt.Fprintln(w, matches["jobid"]+"\t"+job.Name+"\t"+
				job.SubmissionTimeUtc.Local().Format(time.RFC3339)+"\t"+
				job.PrintState)
		}
		w.Flush()
	},
}
