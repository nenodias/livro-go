package main

import (
	"fmt"
)

type Carro struct {
	nome        string
	modelo      string
	ano         int
	combustivel string
}

func (c Carro) ir() {
	fmt.Println("Carro Indo")
}

type Moto struct {
	Carro
}

func main() {
	fmt.Println("Composição")
	carro := Carro{nome: "Gol", modelo: "Ap 1.8", ano: 2010, combustivel: "Flex"}
	moto := Moto{Carro: Carro{nome: "Titan", modelo: "125 KS", ano: 2002, combustivel: "Gasolina"}}
	moto.nome = "Titan"
	fmt.Println(carro)
	fmt.Println(moto)
	carro.ir()
	moto.ir() //Herdado por composição
}
