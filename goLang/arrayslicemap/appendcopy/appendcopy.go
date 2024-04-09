package main

import "fmt"

func main() {

	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array = append(array, 4, 5, 6) // não é possível fazer append em array

	slice1 = append(slice1, 4, 5, 6) // slice é uma referência para um array interno
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	// o copy não redimensiona o slice
	// o copy copia a quantidade de elementos da origem para o destino
	copy(slice2, slice1) // a função copy não redimensiona o slice

	fmt.Println(slice1, slice2)
}
