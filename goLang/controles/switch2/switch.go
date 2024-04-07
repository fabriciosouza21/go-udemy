package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// switch true
	// não é necessário colocar a condição
	// o case é uma expressão que será avaliada
	// o primeiro case que for verdadeiro será executado
	switch {
	case t.Hour() < 12:
		println("Bom dia!")
	case t.Hour() < 18:
		println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
