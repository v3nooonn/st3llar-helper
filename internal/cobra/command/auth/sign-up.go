package auth

import (
	"fmt"

	"github.com/v3nooom/st3llar-helper/internal/cobra/command"

	"github.com/spf13/cobra"
)

// signUp represents the sign-up command
var signUp = &cobra.Command{
	Use:   "sign-up",
	Short: "A brief description of sign-up",
	Long: `A longer description. For example:

sign-up

Cobra is a CLI library for Go that empowers applications.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sign-in called")
	},
}

func init() {
	command.Root.AddCommand(signUp)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aboutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aboutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
