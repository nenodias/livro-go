package exemplos

import (
	"fmt"
	"sort"
)

type Estado struct {
	nome      string
	populacao int
	capital   string
}

func ExemploMapa() {
	//Criando mapaGrande
	fmt.Printf("\n\nMapas\n")
	mapaGrande := make(map[string]string, 4050)
	fmt.Println(mapaGrande)
	capitais := map[string]string{
		"GO": "Goiânia",
	}
	fmt.Println(capitais)

	//Populando mapa de struct
	estados := make(map[string]Estado, 6)
	estados["GO"] = Estado{"Goias", 6434052, "Goiânia"}
	estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
	estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
	estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}

	//Recuperando estado
	sergipe := estados["SE"]
	fmt.Println("Sergipe", sergipe)

	//Chave inválida no mapa
	saoPaulo, encontrado := estados["SP"]
	if encontrado {
		fmt.Println("São Paulo", saoPaulo)
	} else {
		fmt.Println("São Paulo não foi encontrado", saoPaulo)
	}
	//Atualizando valores
	idades := map[string]int{
		"João":    37,
		"Ricardo": 28,
		"Joaquim": 41,
	}
	idades["Joaquim"] = 42
	//Removendo chave
	delete(idades, "João")
	//Exibindo mapa alterado
	fmt.Println(idades)

	//Iterando o mapa
	for sigla, estado := range estados {
		fmt.Println(sigla, estado)
	}

	quadrados := make(map[int]int, 15)
	for n := 1; n <= 15; n++ {
		quadrados[n] = n * n
	}

	numeros := make([]int, 0, len(quadrados))
	fmt.Println("Quadrados", quadrados)
	for n := range quadrados {
		numeros = append(numeros, n)
	}
	//Sem ordem?
	fmt.Println(numeros)

	sort.Ints(numeros)

	//Exibindo os paranaues ordenados
	for _, numero := range numeros {
		quadrado := quadrados[numero]
		fmt.Printf("%d^2 = %d\n", numero, quadrado)
	}
	fmt.Println("FIM")
}
