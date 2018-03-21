// Copyright © 2018 rtslabs

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rtslabs/teamwork-go/configuration"
	"strings"
	"path/filepath"
	"log"
	"github.com/rtslabs/teamwork-go/util"
)

var (
	// config options
	global bool

	// set options
	initConf bool

	// get options
	format string

	// init options
	fileType string
)

// configCmd represents the config sub-command
// no command argument shows help
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get or set various configurations",
	Long:  `Get or set various configurations`,
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the locations of the configs",
	Run: func(cmd *cobra.Command, args []string) {
		for _, conf := range configuration.Configs {
			fmt.Println(conf.Location)
		}
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a config in the current working directory",
	Run: func(cmd *cobra.Command, args []string) {

		fileType = strings.TrimPrefix(fileType, ".")

		path := "."
		if len(args) > 0 {
			path = strings.Join(args, " ")
		}

		dir, err := filepath.Abs(filepath.Dir(path))

		if err != nil {
			log.Fatal("Someting went wrong with your input", path, err)
		}

		if err := configuration.InitConfigDir(dir, fileType); err != nil {
			log.Fatal("Something went wrong while writing config", err)
		}

		fmt.Println("Succesfully created config file at", dir)

	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns the current project config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

var getTeamworkCmd = &cobra.Command{
	Use:   "teamwork",
	Short: "Get teamwork related config",
	Run: func(cmd *cobra.Command, args []string) {

		str, err := util.ToString(configuration.MustGetTeamworkConfig(global), format)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	},
}

var getTeamworkSiteCmd = &cobra.Command{Use: "site", Short: "Get teamwork site name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(configuration.MustGetTeamworkConfig(global).SiteName)
	},
}

var getTeamworkAPIKeyCmd = &cobra.Command{Use: "api-key", Short: "Get teamwork api key",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(configuration.MustGetTeamworkConfig(global).APIKey)
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config parameters to the current working project config",
}

var setTeamworkCmd = &cobra.Command{
	Use:   "teamwork",
	Short: "Set teamwork related config",
}

var setTeamworkSiteCmd = &cobra.Command{Use: "site", Short: "Set teamwork site name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Requires at least one argument")
		}
		conf := mustGetConfig()
		conf.Teamwork.SiteName = args[0]
		configuration.WriteConfig(conf)
	},
}

var setTeamworkAPIKeyCmd = &cobra.Command{Use: "api-key", Short: "Set teamwork api key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Requires at least one argument")
		}
		conf := mustGetConfig()
		conf.Teamwork.APIKey = args[0]
		configuration.WriteConfig(conf)
	},
}

func init() {

	configCmd.AddCommand(configListCmd)

	configCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&fileType, "fileType", "x", "yaml", "Initialize file with given file type extension")

	configCmd.AddCommand(getCmd)
	getCmd.AddCommand(getTeamworkCmd)
	getCmd.PersistentFlags().StringVarP(&format, "format", "f", "json", "Format of output [json, yaml, yml, minified")
	getTeamworkCmd.AddCommand(getTeamworkSiteCmd)
	getTeamworkCmd.AddCommand(getTeamworkAPIKeyCmd)

	configCmd.AddCommand(setCmd)
	setCmd.AddCommand(setTeamworkCmd)
	setTeamworkCmd.AddCommand(setTeamworkSiteCmd)
	setTeamworkCmd.AddCommand(setTeamworkAPIKeyCmd)

	setCmd.PersistentFlags().BoolVarP(&initConf, "init", "i", false, "Initialize a config in the cwd (defaults to false)")

	RootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().BoolVarP(&global, "global", "g", false, "Set or get from the global config in your home directory")
}

func mustGetConfig() (*configuration.Configuration) {
	if global {
		return configuration.MustGetGlobal()
	} else {
		return configuration.MustGetLast()
	}
}
