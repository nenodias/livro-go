package exemplos

import "fmt"

func Exemplos() {
	var exemplo [3]int //Array inicializado de tamanho 3
	//Todas as posições iniciadas com o zero value
	fmt.Println(exemplo)
	exemplo[0], exemplo[1], exemplo[2] = 3, 2, 1
	fmt.Println(exemplo)
	//Exemplo inicializando array com tamanho fixo e valores
	exemplo2 := [5]int{5, 4, 3, 2, 1}
	fmt.Println(exemplo2)
	//Tamanho do array é inferido pelo compilador
	exemplo3 := [...]int{4, 3, 2, 1}
	fmt.Println(exemplo3)
}
