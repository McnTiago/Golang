package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Nome    string    `jason:"name"`
	pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Numeto  int            `json:"entry_number"`
	Especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Nome string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	reponseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(reponseData))

	var responseObject Response
	json.Unmarshal(reponseData, &responseObject)

	fmt.Println(responseObject.Nome)
	fmt.Println(responseObject.pokemon)

	for _, Pokemon := range responseObject.pokemon {
		fmt.Println(Pokemon.Especie.Nome)
	}
}
