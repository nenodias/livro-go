package exemplos

import "fmt"

func ExemploMatriz() {
	//Criando uma matriz de 2x2
	var matriz [2][2]int
	fmt.Println(matriz)
	//Criando uma matriz 2x2 inicializada
	exemplo := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(exemplo)
}
