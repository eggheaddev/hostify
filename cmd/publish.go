package cmd

import (
	"fmt"

	// "os"
	"hostify/connection"
	// "hostify/handlers"
	// "path/filepath"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish your great library",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files := []string{"./cmd", "./connection", "./handlers"}

		fmt.Println("publish package...")

		for file := 0; file < len(files); file++ {
			connection.SendFiles(files[file])
		}
		fmt.Println("Done...")
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
