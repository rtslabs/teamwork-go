package configuration

import (
	"errors"
	"github.com/rtslabs/teamwork-go/util"
	"log"
	"os"
	"time"
)

// config per directory - ordered from home and / to current
var Configs []Configuration

type Configuration struct {
	Location  string
	FileType  string
	Teamwork  TeamworkConfig
	TodoItems []TodoConfig
	Timers    []TimerConfig
	Favorites []FavoriteConfig
	Defaults  FavoriteConfig
}

type TeamworkConfig struct {
	SiteName string
	APIKey   string
	UserId   string
}

type FavoriteConfig struct {
	Name       string
	TaskId     string
	TaskListId string
	ProjectId  string
	Message    string
	Duration   string
	Billable   string
}

type TodoConfig struct {
	Name        string
	TaskId      string
	ProjectId   string
	DueDate     string
	Description string
}

func MustGetLast() *Configuration {
	if len(Configs) == 0 {
		log.Fatal("No configurations found")
	}
	return &Configs[len(Configs)-1]
}

func MustGetGlobal() *Configuration {
	if len(Configs) == 0 {
		log.Fatal("No configurations found")
	}
	return &Configs[0]
}

func InitConfigDir(absPath, extension string) error {
	fileName := absPath + "/" + FILENAME + "." + extension

	if _, err := os.Stat(fileName); err == nil {
		log.Fatal("file " + fileName + " already exists")
	}

	newConfig := Configuration{Location: fileName, FileType: extension}
	return WriteConfig(&newConfig)
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
func MustGetTeamworkConfig(global bool) (config *TeamworkConfig) {

	if global {
		return &MustGetLast().Teamwork
	}

	for i := range Configs {
		conf := &Configs[len(Configs)-i-1]
		if util.NotBlank(conf.Teamwork.SiteName) && util.NotBlank(conf.Teamwork.APIKey) {
			return &conf.Teamwork
		}
	}
	log.Fatal("unable to find valid teamwork config")
	return nil
}
