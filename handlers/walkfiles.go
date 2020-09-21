package handlers

import (
	"fmt"
	"os"
)
// Walkfiles paths
func Walkfiles(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		fmt.Println("directory: "+info.Name(), "path: "+path)
	} else {
		fmt.Println("file: " + path)
	}
	return nil
}
