package handlers

import (
	"fmt"
	"hostify/io"
	"log"
	"os"
)

// InitialPackage create a initial hostify file
func InitialPackage() {

	var path string

	// * get os type
	osType := os.Getenv("OS")

	switch osType {
	case "Windows_NT":
		path = "\\hostify.json"

	// * default path for unix-based os
	default:
		path = "/hostify.json"
	}

	_, errExist := os.Stat(Cwd() + path)

	// * verified if hostify file exist
	if !os.IsNotExist(errExist) {
		io.ErrorMessage("hostify.json file is ready exist")
		os.Exit(1)
	} else {

		file, errorCreate := os.Create("hostify.json")

		if errorCreate != nil {
			io.ErrorMessage("creating hostify.json file\n" + io.Trace)
			log.Fatal(errorCreate)
		}

		bitesWrites, errorWrite := file.WriteString(ManageTemplate())

		if errorWrite == nil {
			file.Close()
			done := fmt.Sprintf("%v bites writes", bitesWrites)
			io.SuccessMessage(done)
			os.Exit(0)
		} else {
			fmt.Println(errorWrite)
			log.Fatal(errorCreate)
		}
	}
}
