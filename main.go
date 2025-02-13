package main

import (
	"time"

	"github.com/vladislav-dd/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
