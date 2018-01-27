package cmd

import (
	"log"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/viper"
	"github.com/swill/teamwork"
)

var (
	conn        *teamwork.Connection
	filter      string
	projects    bool
	assignedTaskSuggestions []prompt.Suggest
)

type task struct {
	taskListName string
	ProjectName  string
	id           int
}

type tasklist struct {
	tasks []task
}

func setAssignedTasks() []prompt.Suggest {
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

	for index := range t {

		task := prompt.Suggest{
			Description: t[index].ProjectName,
			Text:        t[index].TaskListName,
		}

		assignedTaskSuggestions = append(assignedTaskSuggestions, task)
	}

	return assignedTaskSuggestions
}

func GetTasks() []prompt.Suggest {
	return taskSuggestions
}

func InitializeTeamworkData() {
	taskSuggestions = setAssignedTasks()
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
