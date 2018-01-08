package cmd

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"teamworkgo/cmd/localPrompt"
)

// lsCmd represents the ls command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		t := []prompt.Suggest{}
		for l := range t{
			fmt.Printf("%+v\n", l)
		}
	},
}

func init() {
	RootCmd.AddCommand(cdCmd)
	localPrompt.Initialize()
}
