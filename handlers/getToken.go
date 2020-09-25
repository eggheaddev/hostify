package handlers

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
)

// GetToken get token content
func GetToken() string {
	data, err := ioutil.ReadFile(TokenPath())

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

// TokenPath get token path
func TokenPath() string {
	usr, errorGetPath := user.Current()

	if errorGetPath != nil {
		log.Fatal(errorGetPath)
	}

	return filepath.Join(usr.HomeDir, "hostify.key")
}
