package cmd

import (
	"hostify/handlers"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a hostify package",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.InitialPackage()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
