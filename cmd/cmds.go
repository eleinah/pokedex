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
			description: "Exits the program",
			callback:    Exit,
		},
	}
}
