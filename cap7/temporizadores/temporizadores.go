package temporizadores

import (
	"fmt"
	"time"
)

func Executar(c chan<- bool) {
	time.Sleep(5 * time.Second)
	c <- true
}

func Exemplo() {
	fmt.Println("")
	fmt.Println("Exemplo Temporizadores")
	c := make(chan bool, 1)
	go Executar(c)
	fmt.Println("Esperando...")
	fim := false
	for !fim {
		select {
		case fim = <-c:
			fmt.Println("Fim")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout")
			fim = true
		}
	}

	fmt.Println("")
}
