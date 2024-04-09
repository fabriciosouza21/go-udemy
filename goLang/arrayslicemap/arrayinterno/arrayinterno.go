package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2)

	// slice é um ponteiro para um array interno
	// s1 e s2 apontam para o mesmo array interno
	// se o array interno mudar, os dois slices mudam
	s1[0] = 7
	fmt.Println(s1, s2)
}
