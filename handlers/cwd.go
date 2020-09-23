package handlers

import (
	"log"
	"os"
)

// Cwd return currend directory path
func Cwd() string {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}
