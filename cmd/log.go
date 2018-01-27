package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Suggestions ...
	Suggestions bool
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "tasks [filter]",
	Long: `list tasks assigned to a current user
		filter with args`,
	Run: func(cmd *cobra.Command, args []string) {
		if Suggestions {
			t := GetTasks()
			for l := range t {
				fmt.Printf("%+v\n", l)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(logCmd)

	logCmd.PersistentFlags().BoolVarP(&Suggestions, "suggestions", "s", false, "Return prompt suggestions")
}
