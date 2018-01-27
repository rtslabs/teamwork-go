package cmd

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

// completerCmd represents the completer command
var completerCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("this-localPrompt\n")
		fmt.Println("Please use `exit` or `Ctrl-D` to exit this program..")
		t := prompt.New(
			Executor,
			Completer,
			prompt.OptionTitle("tw localPrompt"),
			prompt.OptionPrefix(">>> "),
			prompt.OptionInputTextColor(prompt.Yellow),
		)
		t.Run()
	},
}

func init() {
	RootCmd.AddCommand(completerCmd)
}
