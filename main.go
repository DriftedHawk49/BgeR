package main

import (
	"fmt"
	"log"

	"github.com/BgeR/globals"
	"github.com/BgeR/utilities"
)

func main() {
	/*
		*	Logic
			1. Check for arguements, only 2 arguements are valid, folder containing files, and shuffle time. use golang flags
		*	2. Identify OS type. If not linux, return
		*	3. Identify Desktop Environment -- tested on gnome
			4. check whether the arg is a valid directory.
			5. List out the files present in that, put that in a list
			6. Every minute check for changes in the folders, if change in number of files or any already present file size, update the list
			7. check extension of every file, if the extension is not .png or .jpeg, skip that file.
			8. change the background after every "shuffle-time"
	*/

	flags := utilities.GetFlags()

	osType := utilities.ValidateOSType()
	if osType == globals.OS_INVALID {
		log.Fatal("cannot identify operating system or your operating system is not supported yet")
	}

	de, err := utilities.GetDE()
	if err != nil {
		log.Fatalf("cannot identify Desktop Environment, err : %s", err.Error())
	}

	fmt.Println("flags provided", flags)

	utilities.Init(flags, osType, de)

}
