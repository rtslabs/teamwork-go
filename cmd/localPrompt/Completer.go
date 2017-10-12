package localPrompt

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

const (
	Root = iota
	Projects
	Tasklists
	Tasks
	Task
)

var projectsSuggestions = []prompt.Suggest{}
var taskListsSuggestions = []prompt.Suggest{}
var taskSuggestions = []prompt.Suggest{}

var State = Root
var lsSuggestions = []prompt.Suggest{}
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
	{Text: "exit", Description: "Exit the localPrompt"},
}

func Completer(d prompt.Document) []prompt.Suggest {
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
	case "ls":
		second := args[1]
		if len(args) == 2 {
			subcommands := projectsSuggestions
			return prompt.FilterHasPrefix(subcommands, second, true)
		}
	default:
		return []prompt.Suggest{}
	}

	//switch first {
	//case "ls", "list":
	//	second := args[1]
	//	if len(args) == 2 {
	//	case "projects":
	//		subcommands := projectsSuggestions
	//		return localPrompt.FilterContains(subcommands, first, true)
	//		case "tasklists", "lists":
	//		subcommands := taskListsSuggestions
	//		return localPrompt.FilterHasPrefix(subcommands, first, true)
	//		case "tasks", "tsks":
	//		subcommands := taskSuggestions
	//		return localPrompt.FilterHasPrefix(subcommands, first, true)
	//	}
	//	subcommands := getState()
	//	//printState(State)
	//	return localPrompt.FilterContains(subcommands, first, true)
	//case "cd":
	//	subcommands := getState()
	//	//increaseState(State)
	//	return localPrompt.FilterContains(subcommands, first, true)
	//case "cd ../":
	//	subcommands := getState()
	//	//decreaseState(State)
	//	return localPrompt.FilterContains(subcommands, first, true)
	//
	//}

	return []prompt.Suggest{}
}
