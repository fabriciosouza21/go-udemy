package main

import "fmt"

// contexto lexico é o escopo de uma função
// é o local onde a função foi declarada
// closure é uma função que referencia variáveis de um contexto léxico
func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)
	imprimeX := closure()
	// imprime 10, pois a função closure() referencia a variável x do contexto léxico onde foi declarada
	imprimeX()
}
