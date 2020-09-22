package handlers

import (
	"bufio"
	"fmt"
	"os"
)

// ManageTemplate get the json template
func ManageTemplate() string {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	config := []string{"Package name: ",
		"description: ", "version: ", "entry file: ", "repository: ", "files: "}

	// * get initial config
	for i := 0; i <= len(config)-1; i++ {
		fmt.Printf(config[i])
		scanner.Scan()
		input = append(input, scanner.Text())
	}

	template := fmt.Sprintf(`{
	"name": "%v",
	"description": "%v",
	"version": "%v",
	"entry": "%v",
	"repository": "%v",
	"files": [%v]
}`, input[0],
		input[1],
		input[2],
		input[3],
		input[4],
		input[5])

	return template
}
