package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PokemonMove struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonMoveContainer struct {
	Move PokemonMove `json:"move"`
}

type Pokemon struct {
	Id        int                    `json:"id"`
	Name      string                 `json:"name"`
	Height    int                    `json:"height"`
	Weight    int                    `json:"weight"`
	Order     int                    `json:"order"`
	IsDefault bool                   `json:"is_default"`
	Moves     []PokemonMoveContainer `json:"moves"`
}

func GetPokemon(id int) Pokemon {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	p := Pokemon{}
	err = json.NewDecoder(res.Body).Decode(&p)
	if err != nil {
		panic(err)
	}
	return p
}

func main() {
	pokemons := make(map[int]Pokemon, 150)
	c := make(chan Pokemon)
	for i := 1; i <= 150; i++ {
		go func(n int) {
			c <- GetPokemon(n)
		}(i)
	}
	var done = false
	for {
		select {
		case p := <-c:
			pokemons[p.Id] = p
		case <-time.After(5 * time.Second):
			done = true
		}
		if done {
			break
		}
	}
	for i := 1; i <= 150; i++ {
		fmt.Println(pokemons[i].Name)
	}
	fmt.Println("Fim")
}
