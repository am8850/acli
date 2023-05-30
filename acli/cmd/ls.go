package cmd

import (
	"acli/pkg/repo"
	"log"

	"github.com/spf13/cobra"
)

var (
	lsCmd = &cobra.Command{
		Use:     "ls",
		Short:   "ls list templates",
		Long:    ``,
		Aliases: []string{"list"},
		Run:     runLs,
	}
)

func init() {
}

func findTemplates(path string) []string {
	var list []string
	return list
}

func runLs(ccmd *cobra.Command, args []string) {
	path := "./temp"
	err := repo.CloneToFilesystem(path, "")
	if err != nil {
		log.Fatalln("Error during git clone: ", err)
	}
	templates := findTemplates(path)
	println("Templates:\n", templates)
}
