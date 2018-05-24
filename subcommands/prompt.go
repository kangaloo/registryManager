package subcommands

import "github.com/c-bata/go-prompt"

func completer(d prompt.Document) []prompt.Suggest {

	s := []prompt.Suggest{
		{Text: "show", Description: ""},
		{Text: "set", Description: ""},
		{Text: "exit", Description: ""},
		{Text: "quit", Description: ""},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
