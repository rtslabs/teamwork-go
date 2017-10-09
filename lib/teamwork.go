package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Projects struct {
	ExecutionTime   string    `json:"executionTime"`
	ProjectBeanList []Project `json:"projects"`
}

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TaskLists struct {
	ExecutionTime   string    `json:"executionTime"`
	ProjectBeanList []Task `json:"tasklists"`
}

type Task struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetTask(tasklistId string) (*Task, error) {
	resp := GetRequest("/tasklists/" + tasklistId + "/tasks.json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var s = new(Task)
	err = json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}

	return s, err
}

func GetTaskLists(projectId string) (*TaskLists, error) {
	resp := GetRequest("prjects/" + projectId + "/tasklists.json")

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

func GetProject(projectId string) (*Project, error) {
	resp := GetRequest("prjects/" + projectId)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var s = new(Project)
	err = json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}

	return s, err
}

func GetAllProjects() *Projects {

	resp := GetRequest("projects.json")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s, err := unpackProjects([]byte(body))

	log.Println(s.ProjectBeanList[0].Name)
	log.Println(s.ProjectBeanList[0].Id)

	return s
}

func unpackTasks(body []byte) (*TaskLists, error) {
	var s = new(TaskLists)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}
	return s, err
}

func unpackProjects(body []byte) (*Projects, error) {
	var s = new(Projects)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}
	return s, err
}

