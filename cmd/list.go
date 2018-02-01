package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/swill/teamwork"
)

var (
	// Filter ...
	filterFlag string
	projectsFlag bool
	tasksFlag bool
)


var listCmd = &cobra.Command{
	Use:   "list",
	Short: "tasks [filter]",
	Long: `list tasks assigned to a current user
		filter with args`,
	Run: func(cmd *cobra.Command, args []string) {
		if tasksFlag {
			listTasks()
		} else if projectsFlag {
			listProjects()
		} else {
			listAll()
		}
	},
}

func listTasks() {
	conn := TeamworkConnection()

	// get all projects
	task_ops := &teamwork.GetTasksOps{
	}
	tasks, _, err := conn.GetTasks(task_ops)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	for i := range tasks {
		c := tasks[i].Content
		id := tasks[i].ID

		if (filterText(c)) {
			fmt.Printf("%v\t%v\n", id, c)
		}
	}
}

func listProjects() {
	conn := TeamworkConnection()

	// get all projects
	projects_ops := &teamwork.GetProjectsOps{
		Status: "ALL",
	}
	projects, _, err := conn.GetProjects(projects_ops)
	if err != nil {
		fmt.Printf("Error getting Projects: %s", err.Error())
	}

	for i := range projects {
		p := projects[i].Name
		id := projects[i].ID

		if (filterText(p)) {
			fmt.Printf("%v\t%v\n", id, p)
		}
	}
}

func listAll() {
	// use the list of assigned tasks at init
	for _, i := range tasks {

		projectFilter := filterText(i.ProjectName)
		taskListFilter := filterText(i.TaskListName)
		contentFilter := filterText(i.Content)

		if (projectFilter || taskListFilter || contentFilter) {
			fmt.Printf("ID:		%d\n", i.ID)
			fmt.Printf("Project:	%s\n", i.ProjectName)
			fmt.Printf("Tasklist:	%s\n", i.TaskListName)
			fmt.Printf("Task:		%s\n", i.Content)
			fmt.Print("\n")
		}
	}
}

func filterText(s string) bool {
	substr := strings.TrimSpace(strings.ToUpper(filterFlag))

	if (strings.Contains(s, substr)) {
		return true
	}

	return false
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringVarP(&filterFlag, "filter", "f", "", "Filter responses")
	listCmd.PersistentFlags().BoolVarP(&projectsFlag, "projecs", "p", false, "List Projects")
	listCmd.PersistentFlags().BoolVarP(&tasksFlag, "tasks", "t", false, "List Tasks")
}
