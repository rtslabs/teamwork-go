package cmd

import (
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/swill/teamwork"
)

var taskSuggestions = []prompt.Suggest{}
var tasks = teamwork.Tasks{}
var commands = []prompt.Suggest{
	{Text: "log", Description: "List assigned tasks"},

	// aliases
	// {Text: "list"},

	// customized
	{Text: "bash", Description: "Drop to a bash subshell"},
	{Text: "exit", Description: "Exit the localPrompt"},
}

// Completer ...
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
	case "log":
		second := args[1]
		if len(args) == 2 {
			subCommands := taskSuggestions
			return prompt.FilterHasPrefix(subCommands, second, true)
		}
	default:
		return []prompt.Suggest{}
	}

	return []prompt.Suggest{}
}
