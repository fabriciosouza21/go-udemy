package main

import "fmt"

// as funções são de primeira classe
// ou seja, podem ser atribuidas a variaveis
// e passadas como argumentos

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))

}
