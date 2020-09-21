package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show hostify version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
