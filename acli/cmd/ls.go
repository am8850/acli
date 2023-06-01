package cmd

import (
	"acli/pkg/repo"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	lsCmd = &cobra.Command{
		Use:                   "ls",
		Short:                 "list templates",
		Long:                  ``,
		Aliases:               []string{"list"},
		Run:                   runLs,
		DisableFlagsInUseLine: true,
	}
)

func init() {
}

func listFoldersInFolder(path string) ([]string, error) {
	folders := []string{}

	items, err := os.ReadDir(path + "/templates/")
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.IsDir() {
			folders = append(folders, item.Name())
		}
	}

	return folders, nil
}

func runLs(ccmd *cobra.Command, args []string) {
	err := repo.CloneToFilesystem("", "")
	if err != nil {
		log.Fatalln("Error during git clone: ", err)
	}
	list, err := listFoldersInFolder(repo.Temp_folder_name)
	if err != nil {
		fmt.Println("Unable to get the templates.")
	} else {
		fmt.Println("Templates:")
		for _, folder := range list {
			fmt.Println("  " + folder)
		}
	}
	repo.RemoveTempFolder()
}
