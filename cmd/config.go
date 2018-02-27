// Copyright © 2018 rtslabs

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	initConf bool

	setName       string
	setTaskID     int
	setTaskListID int
	setProjectID  int
	setDate       string
	setMessage    string
	setHours      int
	setMinutes    int
	setBillable   bool
)

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

// addCmd represents the add command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns the current project config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "/shrug I don't know what this command does yet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

// initCmd represents the add command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a config in the current working directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

// setCmd represents the add command
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
	configCmd.AddCommand(addCmd)
	configCmd.AddCommand(initCmd)
	RootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.
	configCmd.PersistentFlags().BoolVarP(&initConf, "init", "i", false, "Initialize a config in the cwd (defaults to false)")

	setCmd.PersistentFlags().StringVarP(&setName, "name", "n", "", "Set an alias [favorite] name to the current working config")
	setCmd.PersistentFlags().IntVarP(&setTaskID, "taskId", "t", 0, "Set a Task ID for the current working config")
	setCmd.PersistentFlags().IntVarP(&setTaskListID, "taskListId", "l", 0, "Set a Task List ID for the current working config")
	setCmd.PersistentFlags().IntVarP(&setProjectID, "projectId", "p", 0, "Set a Project ID for the current working config")
	setCmd.PersistentFlags().StringVarP(&setDate, "date", "d", "", "Set a date to the current working config (mm/dd/yy)")
	setCmd.PersistentFlags().StringVarP(&setMessage, "msg", "m", "", "Set a message to the current working config")
	setCmd.PersistentFlags().IntVar(&setHours, "hours", 0, "Set default hours to the current working config")
	setCmd.PersistentFlags().IntVar(&setMinutes, "min", 0, "Set default minutes to the current working config")
	setCmd.PersistentFlags().BoolVarP(&setBillable, "billable", "b", true, "Set current working config as billable (defaults to true)")
}
