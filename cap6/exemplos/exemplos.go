package exemplos

import (
	"fmt"
	"os"
)

//Função básica sem parâmetros e sem retornos
func ImprimirVersao() {
	fmt.Println("1.12")
}

//Função com  1 parâmetro e sem retorno
func ImprimirSaudacao(nome string) {
	fmt.Println("Olá", nome)
}

//Dois parâmetros
func ImprimirDados(nome string, idade int) {
	fmt.Printf("%s %d anos\n", nome, idade)
}

//Variáveis do mesmo tipo em sequencia
func ImprimirSoma(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func ImprimirSomaTexto(a, b int, texto string) {
	fmt.Printf("%s: %d + %d = %d\n", texto, a, b, a+b)
}

func Somar(a, b int) int {
	return a + b
}

//Múltiplos Retornos
func PrecoFinal(precoCusto float64) (float64, float64) {
	fatorLucro := 1.33
	taxaConversao := 2.34

	precoFinalDollar := precoCusto * fatorLucro
	return precoFinalDollar, precoFinalDollar * taxaConversao
}

//Retorno Nomeado
func PrecoFinalRetorno(precoCusto float64) (
	precoDolar float64, precoReal float64) {
	fatorLucro := 1.33
	taxaConversao := 2.34

	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao
	//return //Pode se omitir os dados no return
	//pois as próprias variaveis definidas serão retornadas
	return precoDolar, precoReal //Mas ainda sim é recomendado escrever o retorno com as vars
}

//
func CriaArquivos(dirBase string, arquivos ...string) {
	for _, nome := range arquivos {
		caminhoArquivo := fmt.Sprintf("%s/%s%s", dirBase, nome, ".txt")
		arq, err := os.Create(caminhoArquivo)
		arq.WriteString(nome)
		defer arq.Close()
		if err != nil {
			fmt.Printf("Erro ao criar o arquivo: %s %v\n", caminhoArquivo, err)
		} else {
			fmt.Printf("Arquivo %s criado\n", caminhoArquivo)
		}
	}
}

func Exemplos() {
	fmt.Println("Início")
	ImprimirVersao()
	ImprimirSaudacao("Usuário")
	ImprimirDados("Cliente", 12)
	ImprimirSoma(2, 4)
	ImprimirSomaTexto(4, 2, "Soma")
	fmt.Printf("%d\n", Somar(3, 4))
	n1, n2 := PrecoFinal(5.0)
	fmt.Printf("%.4f %.4f\n", n1, n2)

	dolar, real := PrecoFinalRetorno(1.0)
	fmt.Printf("Preço final em dólar:%.4f\nPreço final em reais:%.4f\n", dolar, real)

	//Argumentos variáveis
	tmp := os.TempDir()
	CriaArquivos(tmp)
	CriaArquivos(tmp, "teste1")
	CriaArquivos(tmp, "teste2", "teste3", "teste4")
}
