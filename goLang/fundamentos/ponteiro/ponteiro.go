package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros
	// não é possível fazer operações aritméticas com ponteiros

	var p *int = nil

	// & é usado para obter o endereço da variável
	// * é usado para obter o valor da variável
	p = &i // p agora guarda o endereço de i
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiros
	// não é possível fazer operações aritméticas com ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
