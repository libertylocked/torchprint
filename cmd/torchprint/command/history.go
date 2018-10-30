package command

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Aliases: []string{"log"},
	Short:   "View transaction history",
	Run: func(cmd *cobra.Command, args []string) {
		skip, _ := cmd.LocalFlags().GetInt("skip")
		pageSize, _ := cmd.LocalFlags().GetInt("count")
		api := newAPI()

		resp, err := api.GetTransactions(skip, pageSize)
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}
		// pretty print log
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 10, 4, 3, ' ', 0)
		fmt.Fprintln(w, "TIME\tTYPE\tNAME\tDEVICE\tCHARGES")
		for _, tx := range resp.Items {
			fmt.Fprintln(w, tx.Time.Local().Format(time.RFC3339)+"\t"+tx.TransactionType+"\t"+
				tx.JobName+"\t"+tx.Device+"\t"+tx.Amount)
		}
		w.Flush()
	},
}

func init() {
	historyCmd.Flags().Int("skip", 0, "Number of items to skip")
	historyCmd.Flags().Int("count", 20, "Number of items to display on a page")
}
