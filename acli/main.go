package main

import (
	"acli/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln("Error executing acli: ", err)
	}
}
