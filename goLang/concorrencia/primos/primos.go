package main

import (
	"fmt"

	pm "curso.com/m/primos"
)

func main() {
	c := make(chan int, 30)
	go pm.Primos(cap(c), c)
	for i := range c {
		fmt.Printf("%d ", i)
	}

	fmt.Println("Fim!")
}
