package main

import (
	"fmt"
	"time"
)

func main() {

	// funcionar parecido com o while
	i := 1
	for i <= 10 { // não existe while em go
		fmt.Println(i)
		i++
	}

	// for clássico
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")

	// estrutura alinhadas
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}
	}

	// laço infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}

	// Veremos o foreach no capítulo de arrays
}
