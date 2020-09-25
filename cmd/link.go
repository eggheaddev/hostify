package cmd

import (
	"hostify/connection"

	"bufio"
	"fmt"
	"hostify/handlers"
	"hostify/io"
	"os"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link your user key to connect the backend",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// * if a token exist
		if connection.ExistToken() {
			io.WarningMessage(
				fmt.Sprintf("hostify.key is ready exist in Path:\n%v ==> %v %v\n",
					io.Red, io.Reset, handlers.TokenPath()))

			// * over write token message
			fmt.Printf("you want over write the actual token [%vyes %vor%v no%v]: ",
				io.Green, io.Reset, io.Red, io.Reset)
			prompt := bufio.NewScanner(os.Stdin)
			prompt.Scan()
			answer := prompt.Text()

			if answer == "yes" || answer == "y" {
				handlers.LinkKeyHandler()
				connection.ValidateToken()
			} else {
				fmt.Println(io.Red + "\nSkipped...\n" + io.Reset)
			}
		} else {
			handlers.LinkKeyHandler()
			connection.ValidateToken()
		}
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

}
