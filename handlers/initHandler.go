package handlers

import (
	"fmt"
	"hostify/io"
	"os"
)

// InitialPackage create a initial hostify file
func InitialPackage() {

	var path string


	_, errExist := os.Stat(Cwd() + "\\hostify.json")

	// * verified if hostify file exist
	if !os.IsNotExist(errExist) {
		io.ErrorMessage("hostify.json file is ready exist")
		os.Exit(1)
	} else {

		file, err := os.Create("hostify.json")

		if err != nil {
			io.ErrorMessage("creating hostify.json file")
		}

		bitesWriter, err := file.WriteString(`{
	"name": "Name here....",
	"description": "Name here",
	"version": "1.0.0",
	"entry": "...",
	"repository": "https://github.com/{ owner }/{ repo name }",
	"files": ["...", "...", "..."]
}`)

		if err == nil {
			file.Close()
			done := fmt.Sprintf("Done: %v bites writes", bitesWriter)
			io.SuccessMessage(done)
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
