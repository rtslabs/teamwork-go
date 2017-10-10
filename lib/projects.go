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

func GetProject(projectId string) (*Project, error) {
	resp := GetRequest("projects/" + projectId)

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

	return s
}

func unpackProjects(body []byte) (*Projects, error) {
	var s = new(Projects)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	}
	return s, err
}

