package configuration

import (
	"regexp"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"strings"
	"path/filepath"
	"github.com/mitchellh/go-homedir"
	"github.com/rtslabs/teamwork-go/util"
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

	if fileReader, e := os.Open(file); e == nil {

		viper.SetConfigName(FILENAME)

		for _, match := range fileTypeRegex.FindAllString(file, -1) {
			viper.SetConfigType(match)
		}

		viper.ReadConfig(fileReader)

		if unmarshalErr := viper.Unmarshal(&config); unmarshalErr != nil {
			fmt.Println("Error reading config file", file, err)
			err = unmarshalErr
		}

		fileReader.Close()
	} else {
		err = e
		fmt.Println("Error opening config file", file, err)
	}
	config.Location = file
	return config, err
}

/*Parse populates the viper instance walking through the passed paths, loading any
files found in the directory.  It silently ignores any errors (bad files, etc)*/
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
