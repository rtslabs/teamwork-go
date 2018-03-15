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
	Teamwork  TeamworkConfig
	Favorites []FavoriteConfig
	TodoItems []TodoConfig
}

type TeamworkConfig struct {
	SiteName string
	ApiKey   string
	UserId   string
}

type FavoriteConfig struct {
	Name       string
	TaskId     string
	TaskListId string
	ProjectId  string
	Message    string
	Hours      int
	Minutes    int
	Billable   bool
}

type TodoConfig struct {
	Name        string
	TaskId      string
	ProjectId   string
	DueDate     string
	Description string
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
	return TeamworkConfig{}, errors.New("unable to find valid teamwork config")
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
