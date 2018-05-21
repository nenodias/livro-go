package goroutinas

import (
	"fmt"
	"time"
)

func Dormir() {
	fmt.Println("Goroutine dormindo por 5 segundos...")
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine finalizada!")
}

func Imprimir(n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%d", n)
		time.Sleep(200 * time.Millisecond)
	}
}

func Exemplo() {
	fmt.Println("Exemplo")
	go Imprimir(2)
	Imprimir(3)

	fmt.Println("Exemplo 2")
	go Dormir()
	time.Sleep(3 * time.Second)
	fmt.Println("main finalizada")

}
