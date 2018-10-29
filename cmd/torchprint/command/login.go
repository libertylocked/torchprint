package command

import (
	"bufio"
	"os"
	"path"
	"strings"

	"github.com/libertylocked/torchprint"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into printing with username and password",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("username")
		pass, _ := cmd.Flags().GetString("password")
		if len(user) == 0 {
			reader := bufio.NewReader(os.Stdin)
			// ask username in stdin
			cmd.Printf("Enter username: ")
			user, _ = reader.ReadString('\n')
			user = strings.TrimSpace(user)
		}
		if len(pass) == 0 {
			// ask password in stdin
			cmd.Printf("Enter password: ")
			passbytes, _ := gopass.GetPasswd()
			pass = string(passbytes)
		}
		api := (&torchprint.API{}).SetUserPass(user, pass)
		resp, token, err := api.Logon()
		if err != nil {
			cmd.Println(err)
			os.Exit(1)
		}
		userid := resp.Identifier

		// save to config
		viper.Set("userid", userid)
		viper.Set("token", token)
		home, _ := homedir.Dir()
		configFolder := path.Join(home, ".config", "torchprint")
		os.MkdirAll(configFolder, os.ModePerm)
		err = viper.WriteConfigAs(path.Join(configFolder, ".torchprint.json"))
		if err != nil {
			cmd.Println("warning: failed to write config!", err)
		}

		// print id to screen
		cmd.Println("success: userid", userid, "balance", resp.Balance.Amount)
	},
}
