package main

import (
	"acli/cmd"
	"log"
)

func main() {
	//services.CheckArgs("<url>", "<directory>")
	//url := os.Args[1]
	//directory := os.Args[2]
	//services.Clone(url, directory)

	// acli init project-name --template vite-go --git-init
	// acli templates

	err := cmd.Execute()
	if err != nil {
		log.Fatalln("Error executing acli: ", err)
	}
}
