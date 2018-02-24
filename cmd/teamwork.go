// Copyright © 2018 rtslabs


package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// teamworkCmd represents the teamwork command
var teamworkCmd = &cobra.Command{
	Use:   "teamwork",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("teamwork called")
	},
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "tasks [filter]",
	Long: `list tasks assigned to a current user
		filter with args`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("log called")
	},
}

func init() {
	teamworkCmd.AddCommand(logCmd)
	RootCmd.AddCommand(teamworkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// teamworkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// teamworkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
