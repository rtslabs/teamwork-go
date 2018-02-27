package configuration

import (
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
	"fmt"
	"os"
	"path/filepath"
	"github.com/rtslabs/teamwork-go/util"
)

const FILENAME = ".teamworkgo"

func InitConfig(override string) {

	fmt.Println("in init config")
	configDirs := getConfigDirs()

	// read configs
	viper.SetConfigName(FILENAME)
	for e := range configDirs {
		viper.AddConfigPath(configDirs[e])
		viper.MergeInConfig()
	}

	// Use config file from the flag.
	if override != "" {
		viper.SetConfigFile(override)
		viper.MergeInConfig()
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	//siteName := viper.GetString("global.companyName")
	//apiURL := "https://" + siteName + ".teamwork.com/"
	//apiKey := viper.GetString("global.apiKey")

}

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
