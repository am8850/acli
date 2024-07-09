package cmd

import (
	"acli/pkg/repo"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	infoCmd = &cobra.Command{
		Use:                   "info",
		Short:                 "template [template-name]",
		Long:                  ``,
		Aliases:               []string{"information", "Info"},
		Run:                   runInfo,
		DisableFlagsInUseLine: true,
	}
)

func init() {
}

func runInfo(ccmd *cobra.Command, args []string) {
	if len(args) > 0 {
		selectedTemplate := args[0]

		err := repo.CloneToFilesystem("", "")
		if err != nil {
			log.Fatalln("Error during git clone: ", err)
		}

		bytes, _ := os.ReadFile(repo.Temp_folder_name + "/templates/templates.json")
		var templates []struct {
			Id          int    `json:"id"`
			Name        string `json:"name"`
			Path        string `json:"path"`
			Description string `json:"description"`
		}

		err = json.Unmarshal(bytes, &templates)
		if err != nil {
			fmt.Println("Unable to get the templates.")
		}

		if len(templates) > 0 {
			if len(templates) > 0 {
				for _, template := range templates {
					if template.Name == selectedTemplate {
						println("Template: " + template.Name)
						println("Description: " + template.Description)
					}
				}
			}
		}
		repo.RemoveTempFolder()
	}
}
