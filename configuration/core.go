package configuration

import (
	"errors"
	"github.com/rtslabs/teamwork-go/util"
	"log"
	"os"
)

// config per directory - ordered from / to current
var Configs []Configuration

type Configuration struct {
	Location  string
	FileType  string
	Teamwork  TeamworkConfig
	TodoItems []TodoConfig
	Favorites []FavoriteConfig
	Defaults  FavoriteConfig
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
	Time       string
	Billable   string
}

type TodoConfig struct {
	Name        string
	TaskId      string
	ProjectId   string
	DueDate     string
	Description string
}

func InitConfigDir(absPath, extension string) error {
	fileName := absPath + "/" + FILENAME + "." + extension

	if _, err := os.Stat(fileName); err != nil {
		log.Fatal("file " + fileName + " already exists")
	}

	newConfig := Configuration{Location: fileName, FileType: extension}
	return writeConfig(&newConfig)
}

// return todos from all configs
func GetFullTodoList() (todos []TodoConfig) {
	for _, config := range Configs {
		todos = append(todos, config.TodoItems...)
	}
	return todos
}

// return favorite config object found by name
func GetFavorite(name string) (favorite FavoriteConfig, err error) {

	found := false
	for _, config := range Configs {
		util.Overwrite(&config.Defaults, &favorite)
	}
	for _, config := range Configs {
		for _, fav := range config.Favorites {
			if name == fav.Name {
				util.Overwrite(&fav, &favorite)
				found = true
			}
		}
	}

	if !found {
		err = errors.New("favorite not found")
	}

	return favorite, err
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
