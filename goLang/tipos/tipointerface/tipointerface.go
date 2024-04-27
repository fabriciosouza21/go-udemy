package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	// interface vazia é o tipo de todos os tipos, um tipo genérico
	var coisa interface{}

	fmt.Println(coisa)

	coisa = 3

	fmt.Println(coisa)

	// dinamicamente, a interface pode receber qualquer tipo
	type dinamico interface{}

	var coisa2 dinamico = "opa"

	fmt.Println(coisa2)

	coisa2 = true

	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}

	fmt.Println(coisa2)

}
