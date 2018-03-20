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

const FILENAME = ".teamworkgo"

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

func WriteConfig(config *Configuration) (err error) {

	// change extension and delete the old one
	if !strings.HasSuffix(config.Location, config.FileType) {
		os.Remove(config.Location)
		current := filepath.Ext(config.Location)
		naked := strings.TrimSuffix(config.Location, current)
		config.Location = naked + "." + config.FileType
	}

	var data []byte
	switch config.FileType {
	case "json":
		data, err = json.MarshalIndent(config, "", "  ")
	case "yml", "yaml":
		data, err = yaml.Marshal(config)
	default:
		err = errors.New("unrecognized file type: " + config.FileType)
	}

	if err == nil {
		err = ioutil.WriteFile(config.Location, data, 0644)
	}

	return err
}

func readConfig(file string) (config Configuration, err error) {

	if fileData, err := ioutil.ReadFile(file); err == nil {

		config.FileType = "json"
		for _, match := range fileTypeRegex.FindAllString(file, -1) {
			config.FileType = strings.ToLower(match)
		}

		switch config.FileType {
		case "json":
			err = json.Unmarshal(fileData, &config)
		case "yml", "yaml":
			err = yaml.Unmarshal(fileData, &config)
		default:
			err = errors.New("unrecognized file type: " + config.FileType)
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
// ordered from home to current
func getConfigDirs() (dirs []string) {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Unable to find home directory for configs")
		home = "."
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Unable to find working directory for configs")
		path = home
	}

	lastPath := "notarealpath"
	for path != lastPath && path != home {
		dirs = append(dirs, path)
		lastPath = path
		path = filepath.Dir(path)
	}

	dirs = append(dirs, home)
	util.Reverse(dirs)

	return dirs
}
