package handlers

import (
	"fmt"
	"hostify/io"
	"os"
	"runtime"
)

// InitialPackage create a initial hostify file
func InitialPackage() {

	var path string

	_os := runtime.GOOS

	switch _os {
	case "windows":
		path = "\\hostify.json"
	case "darwin":
		path = "/hostify.json"
	case "linux":
		path = "/hostify.json"
	default:
		path = "/hostify.json"
	}

	_, errExist := os.Stat(Cwd() + path)

	// * verified if hostify file exist
	if !os.IsNotExist(errExist) {
		io.ErrorMessage("hostify.json file is ready exist")
		os.Exit(1)
	} else {

		file, err := os.Create("hostify.json")

		if err != nil {
			io.ErrorMessage("creating hostify.json file")
		}

		bitesWriter, err := file.WriteString(ManageTemplate())

		if err == nil {
			file.Close()
			done := fmt.Sprintf("%v bites writes", bitesWriter)
			io.SuccessMessage(done)
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
