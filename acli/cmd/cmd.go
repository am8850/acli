package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	acliCmd = &cobra.Command{
		Use:           "acli",
		Short:         "acli – command-line tool to scaffold an new project",
		Long:          `acli – command-line tool to scaffold a new project.`,
		Version:       "1.0.0",
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
		viper.SetConfigFile(".jamctl")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file used for acli: ", viper.ConfigFileUsed())
	}
}
