package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := getRandomChance(float64(pokemon.BaseExperience))

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res {
		fmt.Printf("%s  was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func getRandomChance(par float64) bool {
	prob := par / 640
	return rand.Float64() > prob
}
