// Copyright © 2018 rtslabs

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/rtslabs/teamwork-go/configuration"
	"github.com/rtslabs/teamwork-go/util"
)

var (
	// list options
	running bool
	stopped bool
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "",
	Long:  ``,
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of today's timers",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		configuration.UpdateAllTimers()

		var runningLines []string
		var stoppedLines []string

		for _, timer := range configuration.GetAllTimers() {
			dur := util.DurationToString(timer.CalculateDuration())

			if timer.IsRunning() {
				runningLines = append(runningLines, fmt.Sprintf("Timer %s has been running for %s", timer.Name, dur))
			} else if !timer.IsRunning() {
				stoppedLines = append(stoppedLines, fmt.Sprintf("Timer %s ran for %s", timer.Name, dur))
			}
		}

		if !running && len(stoppedLines) > 0 {
			fmt.Println("Stopped:")
			util.PrintLines(stoppedLines)
		}

		if !running && len(stoppedLines) > 0 && !stopped && len(runningLines) > 0 {
			fmt.Println() // space between
		}

		if !stopped && len(runningLines) > 0 {
			fmt.Println("Running:")
			util.PrintLines(runningLines)
		}
	},
}

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("switch called")
	},
}

// startCmd represents the switch command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}

// stopCmd represents the switch command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
	},
}

func init() {

	timerCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&running, "running", "r", false, "Display only currently running timers")
	listCmd.Flags().BoolVarP(&stopped, "stopped", "s", false, "Display only currently stopped timers")

	timerCmd.AddCommand(switchCmd)
	timerCmd.AddCommand(startCmd)
	timerCmd.AddCommand(stopCmd)

	RootCmd.AddCommand(timerCmd)
}
