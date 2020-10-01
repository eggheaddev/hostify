package handlers

import (
	"hostify/io"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
)

// GetToken get token content
func GetToken() string {
	data, err := ioutil.ReadFile(TokenPath())

	if err != nil {
		io.ErrorMessage(
			"hostify.key file was not found please add your user token\n" + io.Trace)
		log.Fatal(err)
	}

	return string(data)
}

// TokenPath get token path
func TokenPath() string {
	usr, errorGetPath := user.Current()

	if errorGetPath != nil {
		io.ErrorMessage("getting user home directory\n" + io.Trace)
		log.Fatal(errorGetPath)
	}

	return filepath.Join(usr.HomeDir, "hostify.key")
}
