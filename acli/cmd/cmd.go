package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	Version string = ""

	acliCmd = &cobra.Command{
		Use:           "acli",
		Short:         "acli – command-line tool to scaffold an new project",
		Long:          `acli – command-line tool to scaffold a new project.`,
		Version:       Version,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return acliCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	acliCmd.AddCommand(initCmd)
	acliCmd.AddCommand(lsCmd)
	acliCmd.AddCommand(infoCmd)
	acliCmd.AddCommand(verCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigFile(".acli")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file used for acli: ", viper.ConfigFileUsed())
	}
}
