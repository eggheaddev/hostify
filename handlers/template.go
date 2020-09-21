package handlers

import (
	"fmt"
	"bufio"
	"os"
)

// ManageTemplate get the json template
func ManageTemplate() string {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	config := [6]string {"Package name: ", "description: ","version: ","entry file: ","repository: ", "files: "}
	for i := 0; i <= 5; i++ {
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
		"files": ["...", "...", "..."]
	}`, input[0], input[1], input[2], input[3], input[4])

	return template
}