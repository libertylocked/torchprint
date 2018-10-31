package command

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "torchprint",
	Short: "torchprint is a printjob manager for campus printing at NYU",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var cfgFile string

// Execute runs root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().StringP("username", "u", "", "Login username")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Login password")
	rootCmd.PersistentFlags().String("userid", "",
		"Pharos user ID (if you don't know yours, run the login command)")
	rootCmd.PersistentFlags().String("token", "", "Pharos user token")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("userid", rootCmd.PersistentFlags().Lookup("userid"))

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(historyCmd)
	rootCmd.AddCommand(balanceCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// look for config in $HOME and $HOME/.config/torchprint
		viper.AddConfigPath(path.Join(home, ".config", "torchprint"))
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".torchprint")
	}

	if err := viper.ReadInConfig(); err != nil {
		// config file not found, but dont panic yet
		// since token and userid can be supplied in args
	}
}
