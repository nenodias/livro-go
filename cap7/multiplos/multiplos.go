package multiplos

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Executar(controle *sync.WaitGroup) {
	defer controle.Done()

	duracao := time.Duration(1+rand.Intn(5)) * time.Second
	fmt.Printf("Dormindo por %s...\n", duracao)
	time.Sleep(duracao)
}

func Exemplo() {
	fmt.Println("")
	fmt.Println("Exemplo sincronizando múltiplas gorotinas")

	inicio := time.Now()
	rand.Seed(inicio.UnixNano())

	var controle sync.WaitGroup

	for i := 0; i < 5; i++ {
		controle.Add(1)
		go Executar(&controle)
	}

	controle.Wait()
	fmt.Printf("Finalizado em %s.\n", time.Since(inicio))
	fmt.Println("")
}
