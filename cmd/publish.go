package cmd

import (
	"fmt"
	"hostify/connection"
	"hostify/handlers"
	"path/filepath"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish your great library",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files := []string{"main.go",
			"LICENSE", "go.sum", "go.mod", ".gitignore"}

		fmt.Println("publish package...")
		handlers.ReadJSON()
		connection.SendPackage()
		for file := 0; file < len(files); file++ {

			connection.SendFiles("http://localhost:3000/upload", filepath.Join(handlers.Cwd(), files[file]), "file")
		}
		fmt.Println("Done...")
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
