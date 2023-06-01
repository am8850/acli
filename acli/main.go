package main

import (
	"acli/cmd"
	"fmt"
	"log"

	"github.com/morikuni/aec"
)

const asciiArt = `
             _ _ 
   __ _  ___| (_)
  / _' |/ __| | |
 | (_| | (__| | |
  \__,_|\___|_|_|
`

var (
	version   string
	buildTime string
)

func printASCIIArt() {
	acliLogo := aec.LightGreenF.Apply(asciiArt)
	fmt.Println(acliLogo)
	//fmt.Printf("Version: %s\n\n", version)
}

func main() {

	printASCIIArt()
	cmd.Ver = version
	err := cmd.Execute()
	if err != nil {
		log.Fatalln("Error executing acli: ", err)
	}
}
