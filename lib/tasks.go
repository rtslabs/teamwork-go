package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TaskLists struct {
	ExecutionTime   string     `json:"executionTime"`
	ProjectBeanList []TaskList `json:"tasklists"`
}

type TaskList struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Tasks struct {
	Id           string `json:"id"`
	TaskBeanList []Task `json:"todo-items"`
}

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"content"`
	ProjectName string `json:"project-name"`
}

func GetTasks(tasklistId string) (*Tasks, error) {
	resp := GetRequest("/tasklists/" + tasklistId + "/tasks.json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s, err := unpackTasks([]byte(body))

	return s, err
}

func GetTaskLists(projectId string) (*TaskLists, error) {

	resp := GetRequest("projects/" + projectId + "/tasklists.json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var s = new(TaskLists)
	err = json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}

	return s, err
}

func unpackTasks(body []byte) (*Tasks, error) {
	var s = new(Tasks)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}
	return s, err
}
