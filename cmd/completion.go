package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

// configCmd represents the config command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Print out completion file",
	Long:  `Print out shell file that when sourced creates tab-completion statements for bash`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Should print completion help here")
	},
}

var bashCompletionCmd = &cobra.Command{
	Use:   "bash",
	Short: "Print completion for bash",
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletion(os.Stdout)
	},
}

var zshCompletionCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Print completion for zsh",
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	completionCmd.AddCommand(bashCompletionCmd)
	completionCmd.AddCommand(zshCompletionCmd)

	RootCmd.AddCommand(completionCmd)
}
