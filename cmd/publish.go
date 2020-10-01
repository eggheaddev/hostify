package cmd

import (
	"fmt"
	"hostify/connection"
	"hostify/handlers"
	"strings"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish your great library",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		files := strings.Split(fmt.Sprintf("%v", handlers.ReadJSON()["files"]), " ")

		for i := 0; i < len(files); i++ {
			files[i] = strings.ReplaceAll(files[i], "]", "")
			files[i] = strings.ReplaceAll(files[i], "[", "")
		}

		user := connection.ValidateToken()

		fmt.Println("Publish package...")

		for file := 0; file < len(files); file++ {

			connection.SendFiles(files[file], user, file == (len(files)-1))
		}
		fmt.Println("Done...")

	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
