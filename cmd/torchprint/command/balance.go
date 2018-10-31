package command

import (
	"os"

	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:     "balance",
	Aliases: []string{"bal"},
	Short:   "View account balance",
	Run: func(cmd *cobra.Command, args []string) {
		api := newAPI()
		logon, _, err := api.Logon()
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}
		cmd.Println("Total Balance:", logon.Balance.Amount)
		for _, purse := range logon.Balance.Purses {
			cmd.Println("Purse:", purse.Name, "Amount:", purse.Amount)
		}
	},
}
