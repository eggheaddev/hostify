package cmd

import (
	"hostify/connection"

	"fmt"
	"hostify/handlers"
	"hostify/io"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link your user key to connect the backend",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if connection.ExistToken() {
			io.ErrorMessage(
				fmt.Sprintf("hostify.key is ready exist in Path:\n%v", handlers.TokenPath()))
		} else {
			handlers.LinkKeyHandler()
			connection.ValidateToken()
		}
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

}
