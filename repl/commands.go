package repl

import "github.com/lilleyz7/nutritionCLI/handlers"

type command struct {
	Name        string
	Description string
	Handler     func() error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			Name:        "help",
			Description: "Displays HELP messsage",
			Handler:     handlers.HelpHandler,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit application",
			Handler:     handlers.ExitHandler,
		},
		"add": {
			Name:        "random",
			Description: "get random cocktail recipe",
			Handler:     handlers.AddHandler,
		},
		"search": {
			Name:        "search",
			Description: "get results based on user provided name",
			Handler:     handlers.SearchHandler,
		},
		"view": {
			Name:        "search",
			Description: "get results based on user provided name",
			Handler:     handlers.ViewHandler,
		},
	}
}
