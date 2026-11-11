package main

func main() {
	supportedCommands := buildCommandRegistry()

	firstPageUrl := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	config := config{
		commands:  supportedCommands,
		locations: page{previous: nil, next: &firstPageUrl},
	}
	startRepl(&config)
}
