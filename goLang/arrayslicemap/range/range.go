package main

import "fmt"

func main() {
	// o compilador conta a quantidade de elementos
	// ... é um operador que conta a quantidade de elementos
	numeros := [...]int{1, 2, 3, 4, 5}

	// range retorna o índice e o valor
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// _ ignora o índice
	// range retorna o índice e o valor
	for _, num := range numeros {
		fmt.Println(num)
	}

}
