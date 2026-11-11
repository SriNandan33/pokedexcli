package main

func main() {
	supportedCommands := buildCommandRegistry()
	config := config{
		commands: supportedCommands,
	}
	startRepl(&config)
}
