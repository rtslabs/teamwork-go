package db

import (
	"database/sql"
	"log"
	"strconv"
	"teamworkgo/lib"
)

func PutTaskList(taskList lib.TaskList, project lib.Project, database *sql.DB) {

	tskListId, err := strconv.Atoi(taskList.Id)
	if err != nil {
		log.Fatal(err)
	}

	projId, err := strconv.Atoi(project.Id)
	if err != nil {
		log.Fatal(err)
	}

	// prepare db objects
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS taskLists (id INTEGER PRIMARY KEY, taskListName TEXT, projectId INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	// store db objects
	statement, _ = database.Prepare("INSERT INTO taskLists (id, taskListName, projectId) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(tskListId, taskList.Name, projId)
}

func PutTask(task lib.Task, taskList lib.TaskList, project lib.Project, database *sql.DB) {

	tskListId, err := strconv.Atoi(taskList.Id)
	if err != nil {
		log.Fatal(err)
	}

	projId, err := strconv.Atoi(project.Id)
	if err != nil {
		log.Fatal(err)
	}

	// prepare db objects
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, taskName TEXT, tasklistId INTEGER, projectId INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	// store db objects
	statement, _ = database.Prepare("INSERT INTO tasks (id, taskName, tasklistId, projectId) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(task.Id, task.Name, tskListId, projId)
}
