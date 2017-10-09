package db

import (
	"database/sql"
	"log"
	"strconv"
	"teamworkgo/lib"
)

func PutProject(project lib.Project, database *sql.DB) {
	projId, err := strconv.Atoi(project.Id)
	if err != nil {
		log.Fatal(err)
	}
	// prepare db objects
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS projects (id INTEGER PRIMARY KEY, projectName TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	// store db objects
	statement, err = database.Prepare("INSERT INTO projects (id, projectName) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec(projId, project.Name)
}
