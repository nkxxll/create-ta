package main

import (
	"flag"
	"fmt"
	"os"

	createta "github.com/nkxxll/create-ta"
)

func main() {
	// todo: implement main
	var name string
	var newName string
	var root string
	var help bool
	flag.StringVar(&name, "name", "hello_world", "Name of the TA")
	flag.StringVar(&newName, "newname", "", "The new name of the TA")
	flag.StringVar(&root, "root", ".", "The root of the ta that should be notified")
	flag.BoolVar(&help, "help", false, "Print help menu")

	flag.Parse()

	if help == true {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if newName == "" {
		fmt.Println("Error: No name for the TA provided")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// call the creation
	tan := createta.Create(name, newName)
	tan.GenerateNew(root)
}
