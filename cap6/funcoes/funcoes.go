package funcoes

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func ExemploRegexp() {
	fmt.Println(" ")
	texto := "Anderson tem 21 anos"
	expr := regexp.MustCompile("\\d")
	fmt.Println(expr.ReplaceAllString(texto, "3"))
	fmt.Println(" ")
}

func ExemploRegexpFn() {
	expr := regexp.MustCompile("\\b\\w")
	texto := "antonio carlos jobim"

	//Função anônima
	transformadora := func(s string) string {
		return strings.ToUpper(s)
	}

	processado := expr.ReplaceAllStringFunc(texto, transformadora)
	fmt.Println(processado)
}

func GerarFibonacci(n int) func() {
	return func() {
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b
			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}

func CronometrarFuncao(funcao func()) {
	inicio := time.Now()

	funcao()

	fmt.Printf("\nTempo de execução: %s\n\n", time.Since(inicio))
}

type Agregadora func(n, m int) int

func CalculaSoma(valores []int) int {
	soma := func(n, m int) int {
		return n + m
	}
	return Agregar(valores, 0, soma)
}

func CalculaProduto(valores []int) int {
	soma := func(n, m int) int {
		return n * m
	}
	return Agregar(valores, 1, soma)
}

func Agregar(valores []int, valorInicial int, fn Agregadora) int {
	agregado := valorInicial
	for _, v := range valores {
		agregado = fn(v, agregado)
	}
	return agregado
}

func Exemplos() {
	ExemploRegexp()
	ExemploRegexpFn()
	CronometrarFuncao(GerarFibonacci(8))
	CronometrarFuncao(GerarFibonacci(48))
	CronometrarFuncao(GerarFibonacci(88))
	valores := []int{3, -2, 5, 7, 8, 22, 32, -1}
	fmt.Println(CalculaSoma(valores))
	fmt.Println(CalculaProduto(valores))
}
