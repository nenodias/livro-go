package main

import (
	"fmt"

	"github.com/nenodias/livro-go/cap7/canais"
	"github.com/nenodias/livro-go/cap7/concorrencia"
	"github.com/nenodias/livro-go/cap7/goroutinas"
	"github.com/nenodias/livro-go/cap7/multiplos"
	"github.com/nenodias/livro-go/cap7/temporizadores"
)

func main() {
	fmt.Println("")
	goroutinas.Exemplo()
	fmt.Println("")
	canais.Exemplo()
	temporizadores.Exemplo()
	multiplos.Exemplo()
	concorrencia.Exemplo()
}
