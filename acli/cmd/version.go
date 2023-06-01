package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	verCmd = &cobra.Command{
		Use:                   "ver",
		Short:                 "version information",
		Long:                  ``,
		Aliases:               []string{"version", "Version"},
		Run:                   runVersion,
		DisableFlagsInUseLine: true,
	}
)

var (
	Ver string
)

func runVersion(ccmd *cobra.Command, args []string) {
	fmt.Printf("Version: %s\n", Ver)
}
