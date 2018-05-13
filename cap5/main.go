package main

import (
	"fmt"

	ex "./listaGenerica"
)

type ListaDeCompras []string

func (lista ListaDeCompras) Categorizar() ([]string, []string, []string) {
	var vegetais, carnes, outros []string
	for _, e := range lista {
		switch e {
		case "Alface", "Pepino":
			vegetais = append(vegetais, e)
		case "Atum", "Frango":
			carnes = append(carnes, e)
		default:
			outros = append(outros, e)
		}
	}
	return vegetais, carnes, outros
}

func imprimirSlice(slice []string) {
	fmt.Println("Slice:", slice)
}

func imprimirLista(lista ListaDeCompras) {
	fmt.Println("Lista de Compras:", lista)
}

func main() {
	lista := make(ListaDeCompras, 6)
	lista[0] = "Alface"
	lista[1] = "Pepino"
	lista[2] = "Azeite"
	lista[3] = "Atum"
	lista[4] = "Frango"
	lista[5] = "Chocolate"

	fmt.Println(lista)
	vegetais, carnes, outros := lista.Categorizar()
	fmt.Println("Vegetais", vegetais)
	fmt.Println("Carnes", carnes)
	fmt.Println("Outros", outros)

	slice := []string{"Alface", "Atum", "Azeite"}
	imprimirLista(ListaDeCompras(slice))

	l := ex.ListaGenerica{1, "Café", 42, true, 23, "Bola", 3.14, false}
	fmt.Println("Lista original:", l)
	fmt.Printf("Removendo do início: %v, após a remoção:\n%v\n", l.RemoverInicio(), l)
	fmt.Printf("Removendo do fim: %v, após a remoção:\n%v\n", l.RemoverFim(), l)
	fmt.Printf("Removendo do indice 3: %v, após a remoção:\n%v\n", l.RemoverIndice(3), l)
	fmt.Printf("Removendo do indice 0: %v, após a remoção:\n%v\n", l.RemoverIndice(0), l)
	fmt.Printf("Removendo do último indice: %v, após a remoção:\n%v\n", l.RemoverIndice(len(l)-1), l)
}
