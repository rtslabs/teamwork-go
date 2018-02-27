// Copyright © 2018 rtslabs


package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	initConf		*bool

	addName 		*string
	addTaskID 		*int
	addTaskListID 	*int
	addProjectID 	*int
	addDate 		*string
	addMessage 		*string
	addHours 		*int
	addMinutes 		*int
	addBillable 	*bool
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "I dont know what the parent command does yet /shrug",
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

// setCmd represents the add command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a config in the current project directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		flagKeys := []string{
			"addName",
			"addTaskID",
			"addTaskListID",
			"addProjectID",
			"addDate",
			"addMessage",
			"addHours",
			"addMinutes",
			"addBillable",
		}

		for _, key := range flagKeys {
			//TODO set items to config file
			fmt.Println("Key: {0}", key)
		}

		fmt.Println("add called")
	},
}


func init() {
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(addCmd)
	configCmd.AddCommand(initCmd)
	RootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.
	configCmd.PersistentFlags().BoolVarP(initConf, "init", "i", false, "Initialize a config in the cwd (defaults to false)")

	addCmd.PersistentFlags().StringVarP(addName, "name", "n", "", "Add an alias [favorite] name to the current working config")
	addCmd.PersistentFlags().IntVarP(addTaskID, "taskId", "t", 0,"Add a Task ID for the current working config")
	addCmd.PersistentFlags().IntVarP(addTaskListID, "taskListId", "l", 0, "Add a Task List ID for the current working config")
	addCmd.PersistentFlags().IntVarP(addProjectID, "projectId", "p", 0, "Add a Project ID for the current working config")
	addCmd.PersistentFlags().StringVarP(addDate, "date", "d", "", "Add a date to the current working config (mm/dd/yy)")
	addCmd.PersistentFlags().StringVarP(addMessage, "msg", "m", "", "Add a message to the current working config")
	addCmd.PersistentFlags().IntVarP(addHours, "hours", "h", 0, "Add default hours to the current working config")
	addCmd.PersistentFlags().IntVar(addMinutes, "min", 0, "Add default minutes to the current working config")
	addCmd.PersistentFlags().BoolVarP(addBillable, "billable", "b", true, "Set current working config as billable (defaults to true)")
}
