package cmd

import (
	"acli/pkg/repo"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	template string
	//gitFlag  bool

	initCmd = &cobra.Command{
		Use:     "init",
		Short:   "init [folder_name] --template [template_name]",
		Long:    `acli init - Downloads a template to a target folder.`,
		Aliases: []string{"ini", "Ini", "Init", "create", "Create"},
		Run:     runInit,
	}
)

func runInit(ccmd *cobra.Command, args []string) {
	if len(args) > 0 {
		// 	fmt.Printf("%+v\n", gen.Generate(repo.ReadRepoContents(args[0])))
		// } else {
		// 	fmt.Fprintln(os.Stderr, "No repository is specified. Please specify a valid git repository url.")
		// 	return
		project := args[0]
		msg := fmt.Sprintf("Generating project: %s\nTemplate: %s to ", project, template)
		// if gitFlag {
		// 	msg += "\nWith git init"
		// } else {
		// 	msg += "\nWithout git init"
		// }
		fmt.Println(msg)
		os.Mkdir("./"+project, 0755)
		err := repo.CloneToFilesystem("", "")
		if err != nil {
			fmt.Println("Unable to process template.")
		}
		err = repo.CopyFolder(project, template)
		if err != nil {
			fmt.Println("Unable to process template.")
		}
		repo.RemoveTempFolder()
		fmt.Println("Project created. Instructions:\n  cd " + project + "\n  make install or make install-git")
	} else {
		fmt.Println("No project name is specified for example:\n  acli init my-react-go-app --template react-go")
	}
}

func init() {
	initCmd.PersistentFlags().StringVarP(&template, "template", "t", "react-go", "the name of the template.")
	//initCmd.PersistentFlags().BoolVarP(&gitFlag, "git-init", "g", false, "if the project should be initialized with a git repo.")
	viper.BindPFlag("template", initCmd.PersistentFlags().Lookup("template"))
	//viper.BindPFlag("git-init", initCmd.PersistentFlags().Lookup("git-init"))
}
