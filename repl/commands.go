package repl

import (
	"fmt"

	"github.com/lilleyz7/nutritionCLI/handlers"
)

type command struct {
	Name        string
	Description string
	Handler     func() error
}

func HelpHandler() error {
	commands := getCommands()
	for key, data := range commands {
		fmt.Printf("%s: \n\t %s \n", key, data.Description)
	}
	return nil
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			Name:        "help",
			Description: "Displays HELP messsage",
			Handler:     HelpHandler,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit application",
			Handler:     handlers.ExitHandler,
		},
		"add": {
			Name:        "random",
			Description: "Follow command with a quatity and name of food to add nutrition to intake",
			Handler:     handlers.AddHandler,
		},
		"search": {
			Name:        "search",
			Description: "Follow cammand with a quatity and name of food to get nutrition facts ",
			Handler:     handlers.SearchHandler,
		},
		"view": {
			Name:        "search",
			Description: "View all consumed foods",
			Handler:     handlers.ViewHandler,
		},
		"remove": {
			Name:        "remove",
			Description: "Removes most recent data entry",
			Handler:     handlers.RemoveHandler,
		},
	}
}
