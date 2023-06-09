package commands

import(
	"github.com/krishnanshagarwal112/pokedex/models"
)

func GetCommands() map[string]models.CliCommand {
	return map[string]models.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commandMapB,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore Pokemons at a location",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try catching a pokemon !",
			Callback:    commandCatch,
		},
	}
}