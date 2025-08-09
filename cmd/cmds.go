package cmd

type cliCmd struct {
	name        string
	description string
	callback    func() error
}

func getCmds() map[string]cliCmd {
	return map[string]cliCmd{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    Exit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    Help,
		},
	}
}
