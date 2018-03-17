// Copyright © 2018 rtslabs

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rtslabs/teamwork-go/configuration"
	"strings"
	"path/filepath"
	"log"
)

var (
	initConf bool

	setName       string
	setTaskID     string
	setTaskListID string
	setProjectID  string
	setDate       string
	setMessage    string
	setHours      string
	setMinutes    string
	setBillable   string
)

var fileType string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "/shrug I dont know what the parent command does yet",
	Long: `This could potentially list the current config or print some kind
	of stats for the current project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns the current project config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
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

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config parameters to the current working project config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
	},
}

func init() {
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(setCmd)
	//
	//setCmd.PersistentFlags().StringVarP(&setName, "name", "n", "", "Set an alias [favorite] name to the current working config")
	//setCmd.PersistentFlags().StringVarP(&setTaskID, "taskId", "t", 0, "Set a Task ID for the current working config")
	//setCmd.PersistentFlags().StringVarP(&setTaskListID, "taskListId", "l", 0, "Set a Task List ID for the current working config")
	//setCmd.PersistentFlags().StringVarP(&setProjectID, "projectId", "p", 0, "Set a Project ID for the current working config")
	//setCmd.PersistentFlags().StringVarP(&setDate, "date", "d", "", "Set a date to the current working config (mm/dd/yy)")
	//setCmd.PersistentFlags().StringVarP(&setMessage, "msg", "m", "", "Set a message to the current working config")
	//setCmd.PersistentFlags().StringVarP(&setHours, "hours", 0, "Set default hours to the current working config")
	//setCmd.PersistentFlags().StringVarP(&setBillable, "billable", "b", true, "Set current working config as billable (defaults to true)")

	configCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&fileType, "fileType", "x", "yaml", "Initialize file with given file type extension")

	RootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.
	setCmd.PersistentFlags().BoolVarP(&initConf, "init", "i", false, "Initialize a config in the cwd (defaults to false)")
}
