package main

import (
	"fmt"

	op "github.com/fabriciosouza21/golang-pacotes/operations"
	mfmt "github.com/fabriciosouza21/golang-pacotes/operations/fmt"

	_ "github.com/fabriciosouza21/golang-pacotes/display"
)

func main() {
	fmt.Println("Pacotes")
	r1 := op.Add(1, 2)
	r2 := op.Sub(1, 2)
	fmt.Println(r1, r2)
	mfmt.Println("Olá, mundo!")
	// go run main.go
	// go run *.go
}
