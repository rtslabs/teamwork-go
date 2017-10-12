package localPrompt

import (
	"github.com/c-bata/go-prompt"
	"fmt"
	"strconv"
	"teamworkgo/db"
	"log"
)

func PrintState() []prompt.Suggest {

	t := []prompt.Suggest{}

	switch State {
	case 1:
		t = projectsSuggestions
	default:
		t = projectsSuggestions
	}

	return t
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

func Initialize() {
	// Get the commands from Cobra to build the localPrompt suggestions
	projects := mapProjects()
	for _, project := range projects {
		var promptSuggestion = prompt.Suggest{
			Text:        project.Text,
			Description: project.Description,
		}
		projectsSuggestions = append(projectsSuggestions, promptSuggestion)
	}
}


	func increaseState(state int) prompt.Suggest {
	if state != 3 {
		state = state + 1
	}
	fmt.Println("increasing State to " + strconv.Itoa(state))
	return prompt.Suggest{}
}

func decreaseState(state int) prompt.Suggest {
	if state != 0 {
		state = state -1
	}
	fmt.Println("decreasing State to " + strconv.Itoa(state))
	return prompt.Suggest{}
}

func getState() []prompt.Suggest {
	s := []prompt.Suggest{}

	switch State {
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