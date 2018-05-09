package exemplos

import (
	"fmt"
	"strconv"
)

func ExemploSlice() {
	//Criando um slice
	var a []int
	fmt.Println(a)
	//Criando um slice inicializado
	primos := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primos)
	//Criando slice vazio
	nomes := []string{}
	fmt.Println(nomes)

	//Criando slices com make
	//tipo, tamanho, capacidade do array interno
	nomes2 := make([]string, 10, 20)
	fmt.Println(nomes2)
	for i, _ := range nomes2 {
		nomes2[i] = "nome" + strconv.Itoa(i)
	}
	fmt.Println(nomes2)
	fmt.Println("Tamanho do array", len(nomes2))
	fmt.Println("Capacidade do array interno", cap(nomes2))
	for _, valor := range nomes2 {
		fmt.Println(valor)
	}
	//Exempplo slice com array original
	original := [...]string{"Ana", "Maria", "Letícia", "Horácio", "Lígia"}

	//Criando um slice a partir do indice 1 até o fim
	xerox := original[1:]
	xerox[0] = "Teste"
	fmt.Println("Original", original)
	fmt.Println("Xerox", xerox)

	//Criando um slice a partir do indice 1 até o indice 3 (não inclusivo/sem incluir o indice 3)
	xerox = original[1:3]
	fmt.Println("Xerox", xerox)

	//Slice completo do array original
	xerox = original[:]
	fmt.Println("Xerox", xerox)

	//Adicionando elementos ao slice com append
	xerox = append(xerox, "Eduardo", "José")

	//Mostrando que o array original se manteve inalterado
	fmt.Println("Original", original)

	//O xerox possui os novos elementos
	fmt.Println("Xerox", xerox)

	//Criando uma cópia do slice xerox
	xerox2 := make([]string, len(xerox))
	copy(xerox2, xerox)
	//Exibindo os caras iguaizinhos
	fmt.Println("Xerox", xerox)
	fmt.Println("Xerox2", xerox2)
	//Adicionando item no slice xerox2
	xerox2 = append(xerox2, "Jana")
	//O Slice xerox não foi alterado
	fmt.Println("Xerox", xerox)
	fmt.Println("Xerox2", xerox2)
}
