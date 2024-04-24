package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// não definir o tamanho gerar um slice
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	// passar o slice como parâmetro
	// o operador ... é usado para indicar que o slice deve ser desmembrado em elementos
	// não é possivel passar um array diretamente
	// aprovados := [4]string{"Maria", "Pedro", "Rafael", "Guilherme"} - não funciona
	// aprovados := [...]string{"Maria", "Pedro", "Rafael", "Guilherme"} - não funciona
	imprimirAprovados(aprovados...)
}
