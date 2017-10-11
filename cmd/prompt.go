// Copyright © 2017 Gabriel Duke <gabeduke@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"teamworkgo/db"
	"strings"
	"strconv"
)

const (
	Root = iota
	Projects
	Tasklists
	Tasks
	Task
)


var state = Root
var lsSuggestions = []prompt.Suggest{}
var projectsSuggestions = []prompt.Suggest{}
var taskListsSuggestions = []prompt.Suggest{}
var taskSuggestions = []prompt.Suggest{}
var commands = []prompt.Suggest{
	{Text: "ls", Description: "List the current working tree"},
	{Text: "cd", Description: "change dir"},
	{Text: "cd ../", Description: "change dir up a level"},
	{Text: "projects", Description: "List projects"},
	{Text: "tasklists", Description: "List tasklists"},
	{Text: "tasks", Description: "List tasks"},

	// aliases
	{Text: "list"},
	{Text: "pwd"},
	{Text: "proj"},
	{Text: "tsk"},

	// customized
	{Text: "bash", Description: "Drop to a bash subshell"},
	{Text: "exit", Description: "Exit the prompt"},
}

// completerCmd represents the completer command
var completerCmd = &cobra.Command{
	Use:   "prompt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("this-prompt\n")
		fmt.Println("Please use `exit` or `Ctrl-D` to exit this program..")
		t := prompt.New(
				executor,
				completer,
				prompt.OptionTitle("tw prompt"),
				prompt.OptionPrefix(">>> "),
				prompt.OptionInputTextColor(prompt.Yellow),
				prompt.OptionMaxSuggestion(20),
		)
		t.Run()
	},
}

func executor(t string) {
	if t == "bash" {
		cmd := exec.Command("bash")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	//w := d.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	return argumentsCompleter(args)
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "ls", "list":
		subcommands := getState()
		//printState(state)
		return prompt.FilterContains(subcommands, first, true)
	case "cd":
		subcommands := getState()
		//increaseState(state)
		return prompt.FilterContains(subcommands, first, true)
	case "cd ../":
		subcommands := getState()
		//decreaseState(state)
		return prompt.FilterContains(subcommands, first, true)
	case "projects":
		subcommands := projectsSuggestions
		return prompt.FilterContains(subcommands, first, true)
	case "tasklists", "lists":
		subcommands := taskListsSuggestions
		return prompt.FilterHasPrefix(subcommands, first, true)
	case "tasks", "tsks":
		subcommands := taskSuggestions
		return prompt.FilterHasPrefix(subcommands, first, true)
	default:
		return []prompt.Suggest{}
	}

	return []prompt.Suggest{}
}

func increaseState(state int) prompt.Suggest {
	if state != 3 {
		state = state + 1
	}
	fmt.Println("increasing state to " + strconv.Itoa(state))
	return prompt.Suggest{}
}

func decreaseState(state int) prompt.Suggest {
	if state != 0 {
		state = state -1
	}
	fmt.Println("decreasing state to " + strconv.Itoa(state))
	return prompt.Suggest{}
}

func getState() []prompt.Suggest {
	s := []prompt.Suggest{}

	switch state {
	case 1:
		s = []prompt.Suggest {
			{Text: "projects"},
		}
	case 2:
		s = []prompt.Suggest {
			{Text: "tasklists"},
		}
	case 3:
		s = []prompt.Suggest {
			{Text: "Tasks"},
		}
	case 4:
		s = []prompt.Suggest {
			{Text: "Task"},
		}
	default:
		s = []prompt.Suggest {
			{Text: "Root"},
		}
	}

	return s
}

func printState(state int) prompt.Suggest {
	switch state {
	case 1:
		fmt.Println("%+v", projectsSuggestions)
	default:
		//TODO
		fmt.Println("help")
	}

	return prompt.Suggest{}
}

func mapProjects() []prompt.Suggest {
	rows := db.GetProjects()
	suggestions := []prompt.Suggest{}

	for rows.Next() {
		var r prompt.Suggest
		err := rows.Scan(&r.Description, &r.Text)
		if err != nil {
			log.Fatal(err)
		}
		suggestions = append(suggestions, r)
	}

	return suggestions
}

func initializePrompt() {
	// Get the commands from Cobra to build the prompt suggestions
	projects := mapProjects()
	for _, project := range projects {
		var promptSuggestion = prompt.Suggest{
			Text:        project.Text,
			Description: project.Description,
		}
		projectsSuggestions = append(projectsSuggestions, promptSuggestion)
	}
}

func init() {
	RootCmd.AddCommand(completerCmd)
	initializePrompt()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
