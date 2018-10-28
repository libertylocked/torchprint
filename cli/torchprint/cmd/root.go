package cmd

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
		// Do Stuff Here
		fmt.Println("in root cmd")
	},
}

var cfgFile string

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
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
		viper.AddConfigPath(home)
		viper.AddConfigPath(path.Join(home, ".config", "torchprint"))
		viper.SetConfigType("json")
		viper.SetConfigName(".torchprint.json")
	}

	viper.ReadInConfig()
}
