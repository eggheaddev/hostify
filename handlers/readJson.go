package handlers

import (
	"encoding/json"
	"hostify/io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ReadJSON ....
func ReadJSON() map[string]interface{} {

	var hostify map[string]interface{}

	_, err := os.Stat(filepath.Join(Cwd(), "hostify.json"))

	if os.IsNotExist(err) {
		io.ErrorMessage("the hostify.json file was not found")
		os.Exit(1)
	} else {

		file, err := ioutil.ReadFile(filepath.Join(Cwd(), "hostify.json"))

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(file, &hostify)

	}
	return hostify
}
