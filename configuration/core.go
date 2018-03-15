package configuration

import (
	"errors"
	"github.com/rtslabs/teamwork-go/util"
	"reflect"
)

// config per directory
var Configs []Configuration

type Configuration struct {
	Location  string
	Teamwork  TeamworkConfig   `mapstructure:"teamwork"`
	Favorites []FavoriteConfig `mapstructure:"favorites"`
	TodoItems []TodoConfig     `mapstructure:"todoItems"`
}

type TeamworkConfig struct {
	SiteName string `mapstructure:"siteName"`
	ApiKey   string `mapstructure:"apiKey"`
	UserId   string `mapstructure:"userId"`
}

type FavoriteConfig struct {
	Name       string `mapstructure:"name"`
	TaskId     string `mapstructure:"taskId"`
	TaskListId string `mapstructure:"taskListId"`
	ProjectId  string `mapstructure:"projectId"`
	Message    string `mapstructure:"message"`
	Hours      int    `mapstructure:"hours"`
	Minutes    int    `mapstructure:"minutes"`
	Billable   bool   `mapstructure:"billable"`
}

type TodoConfig struct {
	Name        string `mapstructure:"name"`
	TaskId      string `mapstructure:"taskId"`
	ProjectId   string `mapstructure:"projectId"`
	DueDate     string `mapstructure:"dueDate"`
	Description string `mapstructure:"description"`
}

const FILENAME = ".teamworkgo"

// return todos from all configs
func GetFullTodoList() (todos []TodoConfig) {
	for _, config := range Configs {
		todos = append(todos, config.TodoItems...)
	}
	return todos
}

// return favorite config object found by name
func GetFavorite(name string) (favorite FavoriteConfig, err error) {
	for _, config := range Configs {
		for _, fav := range config.Favorites {
			if fav.Name == name {
				return fav, nil
			}
		}
	}
	return FavoriteConfig{Name: name}, errors.New("Unable to find favorite " + name)
}

// return favorite config object found by name
func GetTeamworkConfig() (config TeamworkConfig, err error) {
	for _, conf := range Configs {
		if util.NotBlank(conf.Teamwork.SiteName) && util.NotBlank(conf.Teamwork.ApiKey) {
			return conf.Teamwork, nil
		}
	}
	return TeamworkConfig{}, errors.New("Unable to find valid teamwork config")
}

// TODO get working
func overwrite(in interface{}, out interface{}) {

	t := reflect.TypeOf(in)

	inPtr := reflect.ValueOf(in)

	out2 := reflect.New(reflect.TypeOf(out))
	outPtr := reflect.ValueOf(out2)

	for i := 0; i < t.NumField(); i++ {

		// Ignore fields that don't have the same type as a string
		if t.Field(i).Type != reflect.TypeOf("") {
			continue
		}

		inField := inPtr.Field(i)
		str := inField.Interface().(string)
		if util.NotBlank(str) {
			outField := outPtr.Field(i)
			outField.SetString(str)
		}
	}
}
