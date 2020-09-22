package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish your great library",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publish package...")
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
