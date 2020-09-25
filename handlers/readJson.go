package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type hostifyJSON struct {
	name        string
	description string
	version     string
	entry       string
	repository  string
	files       []interface{}
}

// ReadJSON ....
func ReadJSON() {

	file, err := ioutil.ReadFile(filepath.Join(Cwd(), "hostify.json"))

	if err != nil {
		fmt.Println(err)
	}

	var hostify map[string]interface{}

	json.Unmarshal(file, &hostify)

	// var data hostifyJSON = hostifyJSON{
	// 	name:  			hostify["name"],  			description: hostify["description"],
	// 	entry: 			hostify["entry"],
	// 	repository: hostify["repository"],	version: hostify["version"],
	// }

	// files := make([]string, int(hostify["files"]))

	fmt.Println("ok")
	fmt.Printf("%T ", hostify["files"])
}
