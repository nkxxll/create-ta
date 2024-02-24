package main

import (
	"flag"
	"fmt"
	"os"

	createta "github.com/nkxxll/create-ta"
	log "github.com/sirupsen/logrus"
)

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if !ok {
		lvl = "debug"
	}
	// parse string, this is built-in feature of logrus
	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.DebugLevel
	}
	// set global log level
	log.SetLevel(ll)
}

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
		fmt.Println("Error: No new name for the TA provided")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Debugf("Args: name: %s, newname: %s, root: %s", name, newName, root)
	// call the creation
	tan := createta.Create(name, newName)
	tan.GenerateNew(root)
}
