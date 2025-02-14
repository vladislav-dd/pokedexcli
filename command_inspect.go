package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	value, exist := cfg.caughtPokemon[name]
	if exist {
		fmt.Printf("Name: %v\n", value.Name)
		fmt.Printf("Height: %v\n", value.Height)
		fmt.Printf("Weight: %v\n", value.Weight)
		fmt.Println("Stats:")
		for _, stat := range value.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range value.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
