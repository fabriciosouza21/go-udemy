package main

import "fmt"

func main() {

	s := make([]int, 10)
	s[9] = 12

	fmt.Println(s)

	// primeiro argumento é o tipo do slice
	// segundo argumento é o tamanho do slice
	// terceiro argumento é tamanho do array interno
	s = make([]int, 10, 20)

	// len() retorna o tamanho do slice
	// cap() retorna a capacidade do slice (tamanho do array interno)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(s, len(s), cap(s))

	// se adicionar mais um elemento, o slice dobra de tamanho
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
