package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 0
	for {
		n++
		i := rand.Intn(4200)
		fmt.Println(i)

		if i%42 == 0 {
			break
		}
	}
	fmt.Println("Saída após %d iterações.\n", n)

	var i int
externo:
	for {
		for i = 0; i < 10; i++ {
			if i == 5 {
				break externo
			}
		}
	}
	fmt.Println(i)

loop:
	for i = 0; i < 10; i++ {
		fmt.Printf("for i = %d\n", i)

		switch i {
		case 2, 3:
			fmt.Printf("Quebrando switch, i == %d\n", i)
			break
		case 5:
			fmt.Printf("Quebrando loop, i == %d\n", i)
			break loop
		}
	}
}
