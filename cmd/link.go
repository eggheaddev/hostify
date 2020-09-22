package cmd

import (
	// "fmt"
	"hostify/handlers"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link your user key to connect the backend",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.LinkKeyHandler()
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

}
