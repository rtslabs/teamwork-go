package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// Filter ...
	Filter string
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "tasks [filter]",
	Long: `list tasks assigned to a current user
		filter with args`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, i := range tasks {
			substr := strings.TrimSpace(strings.ToUpper(Filter))

			t := strings.ToUpper(i.TaskListName)
			d := strings.ToUpper(i.ProjectName)


			taskFilter := strings.Contains(t, substr)
			descriptionFilter := strings.Contains(d, substr)

			if (taskFilter || descriptionFilter) {
				fmt.Printf("ID:		%d\n", i.ID)
				fmt.Printf("Task:		%s\n", i.TaskListName)
				fmt.Printf("Project:	%s\n", i.ProjectName)
				fmt.Print("\n")
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&Filter, "filter", "f", "", "Filter Tasks")
}
