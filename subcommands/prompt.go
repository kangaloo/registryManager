package subcommands

import "github.com/c-bata/go-prompt"

func completer(d prompt.Document) []prompt.Suggest {

	s := []prompt.Suggest{
		{Text: "show", Description: "List all docker insecure registries"},
		{Text: "set", Description: "Reset the docker config path"},
		{Text: "exit", Description: "Quit this manager"},
		{Text: "quit", Description: "Quit this manager"},
		{Text: "help", Description: "Print usage"},
		{Text: "print", Description: "Print the config file"},
		{Text: "pcd", Description: "Print current config dir"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
