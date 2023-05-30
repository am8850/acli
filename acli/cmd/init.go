package cmd

import (
	"acli/pkg/repo"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	template string
	//gitFlag  bool

	initCmd = &cobra.Command{
		Use:   "init project_name",
		Short: "init a project with a template",
		Long:  ``,
		Run:   runInit,
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
		err := repo.CloneToFilesystem("", "")
		if err != nil {
			log.Fatalln("Error during git clone: ", err)
		}

	}
}

func init() {
	initCmd.PersistentFlags().StringVarP(&template, "template", "t", "vite-go", "the name of the template.")
	//initCmd.PersistentFlags().BoolVarP(&gitFlag, "git-init", "g", false, "if the project should be initialized with a git repo.")
	viper.BindPFlag("template", initCmd.PersistentFlags().Lookup("template"))
	//viper.BindPFlag("git-init", initCmd.PersistentFlags().Lookup("git-init"))
}
