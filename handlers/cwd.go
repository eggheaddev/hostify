package handlers

import (
	"fmt"
	"os"
)

// Cwd return currend directory path
func Cwd() string {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	return cwd
}
