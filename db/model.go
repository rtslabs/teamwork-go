package db

import (
	"log"
	"strconv"
	"teamworkgo/lib"
	"database/sql"
)

var dbContext = "/home/gabeduke/.go/src/teamworkgo/twgo.db"

func PrepareProjectsTable() {
	database, _ := Open(dbContext)
	// prepare db objects
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS projects (id INTEGER PRIMARY KEY, projectName TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	database.Close()
}

func PrepareTasksTable() {
	database, _ := Open(dbContext)
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, taskName TEXT, tasklistId INTEGER, projectId INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	database.Close()
}

func PrepareTaskListsTable() {
	database, _ := Open(dbContext)
    statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS taskLists (id INTEGER PRIMARY KEY, taskListName TEXT, projectId INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
    database.Close()
}

func GetProjects() *sql.Rows{
	database, _ := Open(dbContext)
	// query
	rows, err := database.Query("SELECT * FROM projects")
	checkErr(err)

	return rows
}

func PutProject(project lib.Project) {
	database, _ := Open(dbContext)
	projId, err := strconv.Atoi(project.Id)
	if err != nil {
		log.Fatal(err)
	}

	// store db objects
	statement, _ := database.Prepare("INSERT INTO projects (id, projectName) VALUES ( ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

    statement.Exec(projId, project.Name)
	database.Close()
}

func PutTask(task lib.Task, taskList lib.TaskList, project lib.Project) {

	database, _ := Open(dbContext)

	tskListId, err := strconv.Atoi(taskList.Id)
	if err != nil {
		log.Fatal(err)
	}

	projId, err := strconv.Atoi(project.Id)
	if err != nil {
		log.Fatal(err)
	}

	statement, _ := database.Prepare("INSERT INTO tasks (id, taskName, tasklistId, projectId) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(task.Id, task.Name, tskListId, projId)
	database.Close()
}

func PutTaskList(taskList lib.TaskList, project lib.Project) {

	database, _ := Open(dbContext)

	tskListId, err := strconv.Atoi(taskList.Id)
	if err != nil {
		log.Fatal(err)
	}

	projId, err := strconv.Atoi(project.Id)
	if err != nil {
		log.Fatal(err)
	}

	// store db objects
	statement, _ := database.Prepare("INSERT INTO taskLists (id, taskListName, projectId) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(tskListId, taskList.Name, projId)
	database.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}