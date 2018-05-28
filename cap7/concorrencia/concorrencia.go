package concorrencia

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func calcular(base float64, controle *sync.WaitGroup) {
	defer controle.Done()
	n := 0.0

	for i := 0; i < 100000000; i++ {
		n += base / math.Pi * math.Sin(2)
	}

	fmt.Println(n)
}

func Exemplo() {
	//runtime.GOMAXPROCS(runtime.NumCPU())//Teste com o total de cores do pc
	runtime.GOMAXPROCS(1) //Testando com 1 core
	fmt.Println("")
	fmt.Println("Exemplo de Concorrência")
	inicio := time.Now()
	var controle sync.WaitGroup
	controle.Add(3)

	go calcular(9.37, &controle)
	go calcular(6.94, &controle)
	go calcular(42.57, &controle)

	controle.Wait()
	fmt.Printf("Finalizado em %s\n", time.Since(inicio))

	fmt.Println("")
}
