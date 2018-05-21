package canais

import (
	"fmt"
	"time"
)

func Produzir(c chan int) {
	c <- 33
}

func ProduzirBuff(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
}

//Função que envia valores para o canal
func ProduzirFluxo(c chan<- int) {
	c <- 1
	c <- 2
	c <- 3
}

//Função consome valores do channels
func ConsumirFluxo(c <-chan int) {
	for valor := range c {
		fmt.Print(valor)
	}
}

func Separar(nums []int, i, p chan<- int, pronto chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	pronto <- true
}

func Exemplo() {
	fmt.Println("Exemplo")
	c := make(chan int)
	go Produzir(c)
	valor := <-c
	fmt.Println("Valor:", valor)
	buff := make(chan int, 3)
	go ProduzirBuff(buff)
	//fmt.Println(<-buff, <-buff, <-buff, <-buff)
	/*
		for {
			valor, ok := <-buff
			if ok {
				fmt.Println("V", valor)
			} else {
				break
			}
		}
	*/
	for valor := range buff {
		fmt.Println(valor)
	}

	fmt.Println("Funções com Fluxo")
	bufie := make(chan int, 3)
	go ProduzirFluxo(bufie)
	go ConsumirFluxo(bufie)
	time.Sleep(2 * time.Second)

	i, p := make(chan int), make(chan int)
	pronto := make(chan bool)
	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}

	go Separar(nums, i, p, pronto)
	var impares, pares []int
	fim := false
	for !fim {
		select {
		case n := <-i:
			impares = append(impares, n)
		case n := <-p:
			pares = append(pares, n)
		case fim = <-pronto:
		}
	}
	fmt.Printf("Ímpares: %v | Pares: %v\n", impares, pares)
}
