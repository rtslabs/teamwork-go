package configuration

import (
	"regexp"
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"github.com/mitchellh/go-homedir"
	"github.com/rtslabs/teamwork-go/util"
	"encoding/json"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"errors"
)

var fileRegex = regexp.MustCompile(".*\\.teamworkgo\\.(yaml|yml|json)$")
var fileTypeRegex = regexp.MustCompile(`[^.]+$`)

func InitConfig(override string) {

	configDirs := getConfigDirs()

	// read configs
	files := getConfigsFromDirs(configDirs)
	Configs = readFileConfigs(files)

	// Use config file from the flag.
	if override != "" {
		if config, err := readConfig(override); err == nil {
			Configs = append(Configs, config)
		}
	}
}

func readConfig(file string) (config Configuration, err error) {

	if fileData, err := ioutil.ReadFile(file); err == nil {

		fileType := "json"
		for _, match := range fileTypeRegex.FindAllString(file, -1) {
			fileType = strings.ToLower(match)
		}

		switch fileType {
		case "json":
			err = json.Unmarshal(fileData, &config)
		case "yml", "yaml":
			err = yaml.Unmarshal(fileData, &config)
		default:
			err = errors.New("unrecognized file type: " + fileType)
		}

	}

	if err != nil {
		fmt.Println("error opening config file", file, err)
	}

	config.Location = file
	return config, nil
}

func readFileConfigs(files []string) (configs []Configuration) {

	for _, file := range files {
		if config, err := readConfig(file); err == nil {
			configs = append(configs, config)
		}
	}
	return configs
}

// looks for the config files in the given directories
func getConfigsFromDirs(dirs []string) (files []string) {

	for _, pa := range dirs {
		pa = os.ExpandEnv(pa)

		walk := func(p string, i os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !i.IsDir() {
				if fileRegex.MatchString(p) {
					formatted := strings.Replace(p, "\\", "/", -1)
					files = append(files, formatted)
				}

			} else if pa != p {
				return filepath.SkipDir
			}
			return nil
		}
		filepath.Walk(pa, walk)
	}

	return files
}

// gets a list of config directories in the order they should be read
func getConfigDirs() []string {

	var dirs []string
	if path, err := os.Getwd(); err == nil {
		for path != "/" {
			dirs = append(dirs, path)
			path = filepath.Dir(path)
		}
		dirs = append(dirs, "/")
	} else {
		fmt.Println("Unable to find working directory for configs")
	}

	if home, err := homedir.Dir(); err == nil {
		if !util.Contains(dirs, home) {
			dirs = append(dirs, home)
		}
	} else {
		fmt.Println("Unable to find home directory for configs")
	}

	util.Reverse(dirs)

	return dirs
}
