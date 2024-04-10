package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[12345678979] = "Pedro"
	aprovados[12345678980] = "Ana"
	// elementos com a mesma chave são sobrescritos

	fmt.Println(aprovados)

	// 1 - chave
	// 2 - valor
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	// não gera erro
	// imprime uma string vazia
	fmt.Println(aprovados[12345678978])

}
