package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice não é um array! Slice define um pedaço de um array.
	// inclui o primeiro elemento, mas exclui o último
	// o slice não cria um novo array, ele referencia um array já existente
	// o slice é uma estrutura que tem um ponteiro para um array e um tamanho
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2] // novo slice, mas aponta para o mesmo array
	fmt.Println(a2, s3)

	// você pode imaginar um slice como: tamanho e um ponteiro para um array
	// o slice de um slice é um novo slice que aponta para o mesmo array
	s4 := s2[:1]
	fmt.Println(s2, s4)

}
