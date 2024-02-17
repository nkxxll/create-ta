package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
    // todo: implement main
    var name string
    var help bool
    flag.StringVar(&name, "name", "", "Name of the TA")
    flag.BoolVar(&help, "help", false, "Print help menu")

    flag.Parse()
    
    if (help == true) {
        flag.PrintDefaults()
        os.Exit(1)
    }
    if (name == "") {
        fmt.Println("Error: No name for the TA provided")
        flag.PrintDefaults()
        os.Exit(1)
    }

    // call the creation
}
