package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
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

type GenericSyncMap[K comparable, V any] struct {
	sync.Map
}

func (p *GenericSyncMap[K, V]) Put(key K, value V) {
	p.Store(key, value)
}

func (p *GenericSyncMap[K, V]) Get(key K) (V, error) {
	v, loaded := p.Load(key)
	if !loaded {
		return *new(V), errors.New(fmt.Sprintf("NotFoundError: key=%v", key))
	}
	return v.(V), nil
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
	pokemons := GenericSyncMap[int, Pokemon]{}
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
			pokemons.Put(p.Id, p)
		case <-time.After(5 * time.Second):
			done = true
		}
		if done {
			break
		}
	}
	for i := 1; i <= 150; i++ {
		p, _ := pokemons.Get(i)
		fmt.Println(i, p.Name)
	}
	p, err := pokemons.Get(151)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Name)
	fmt.Println("Fim")
}
