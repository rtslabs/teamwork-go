package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/swill/teamwork"
	"github.com/c-bata/go-prompt"
)

type Task struct {
	taskListName string
	ProjectName  string
	id           int
}

type Tasklist struct {
	Tasks []Task
}

var (
	conn     *teamwork.Connection
	filter   string
	projects bool
	MyTasks  Tasklist
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "tasks [filter]",
	Long: `list tasks assigned to a current user
		filter with args`,
	Run: func(cmd *cobra.Command, args []string) {
		t := prompt.Input(">>> ", taskCompleter)
		fmt.Println("You selected " + t)
	},
}

func buildTaskList() []prompt.Suggest {
	conn = tw()
	userID := viper.GetString("global.userId")

	// get all tasks
	taskOps := &teamwork.GetTasksOps{
		ResponsiblePartyIDs: userID,
	}

	t, _, err := conn.GetTasks(taskOps)
	if err != nil {
		log.Fatal(err)
	}

	suggestions := []prompt.Suggest{}

	for index := range t {

		task := prompt.Suggest{
			Description:	t[index].ProjectName,
			Text:			t[index].TaskListName,
		}

		suggestions = append(suggestions, task)
	}

	return suggestions
}

func taskCompleter(d prompt.Document) []prompt.Suggest {
	s := buildTaskList()
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}


func tw() *teamwork.Connection {
	apiToken := viper.GetString("global.apiKey")

	// setup the teamwork connection
	conn, err := teamwork.Connect(apiToken)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return conn
}

func init() {
	RootCmd.AddCommand(logCmd)
	// twCmd.PersistentFlags().BoolVarP(&projects, "projects", "p", false, "list projects")
	logCmd.PersistentFlags().StringVarP(&filter, "filter", "f", "", "filter tasks")

	// twCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
}
