package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type produto struct {
	ID    int      `json:"id", xml:"id"`
	Nome  string   `json:"nome", xml:"nome"`
	Preco float64  `json:"preco", xml:"preco"`
	Tags  []string `json:"tags", xml:"tags"`
}

func main() {
	fmt.Println("Exemplo JSON")
	p1 := produto{1, "Maçã", 3.50, []string{"Fruta", "Natural"}}
	dados, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Erro ao parsear p1 para json")
		os.Exit(1)
	}
	p1Json := string(dados)
	fmt.Println(p1Json)

	var p2 produto
	p2Json := `{"id":2,"nome":"Notebook","preco":1400.50,"tags":["informática", "tecnologia"]}`
	err = json.Unmarshal([]byte(p2Json), &p2)
	if err != nil {
		fmt.Println("Erro ao parsear p2 para struct")
		os.Exit(1)
	}
	fmt.Println(p2)

	dados, err = xml.MarshalIndent(p1, " ", "    ")
	if err != nil {
		fmt.Println("Erro ao parsear p1 para xml")
		os.Exit(1)
	}
	xmlP1 := string(dados)
	fmt.Println(xmlP1)

	var p3 produto
	xmlP3 := `
	<produto>
		<ID>3</ID>
		<Nome>Suco</Nome>
		<Preco>5.25</Preco>
		<Tags>Fruta</Tags>
		<Tags>Natural</Tags>
		<Tags>Bebidas</Tags>
	</produto>`
	err = xml.Unmarshal([]byte(xmlP3), &p3)

	if err != nil {
		fmt.Println("Erro ao parsear p3 para struct")
		os.Exit(1)
	}
	fmt.Println(p3)
}
