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
)

var projectsSuggestions = []prompt.Suggest{}
var taskListsSuggestions = []prompt.Suggest{}
var taskSuggestions = []prompt.Suggest{}
var commands = []prompt.Suggest{
	{Text: "ls", Description: "List the current working tree"},
	{Text: "cd", Description: "change dir"},

	// aliases
	{Text: "list"},

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
		defer fmt.Println("Bye!")
		t := prompt.Input(">>> ", completer, prompt.OptionMaxSuggestion(20))
		executor(t)
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
	return
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

	//s := argumentsCompleter
	//return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	return argumentsCompleter(args)
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "projects", "proj":
		second := args[1]
		subcommands := projectsSuggestions
		return prompt.FilterHasPrefix(subcommands, second, true)
	case "tasklists":
		third := args[2]
		subcommands := taskListsSuggestions
		return prompt.FilterHasPrefix(subcommands, third, true)
	case "tasks":
		fourth := args[3]
		subcommands := taskSuggestions
		return prompt.FilterHasPrefix(subcommands, fourth, true)
	default:
		return []prompt.Suggest{}
	}

	return []prompt.Suggest{}
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
