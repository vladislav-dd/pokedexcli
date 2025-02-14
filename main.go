package main

import (
	"math/rand"
	"time"

	"github.com/vladislav-dd/pokedexcli/internal/pokeapi"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
