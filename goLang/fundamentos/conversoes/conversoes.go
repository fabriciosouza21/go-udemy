package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Go não permite conversões implicitas

	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	// profesor otario
	// converte o valor para baixo
	// 6
	fmt.Println(notaFinal)

	// cuidado...
	// pegar o valor da tabela ASCII
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	// retorna dois valores, primeiro o valor convertido, segundo um erro
	num, _ := strconv.Atoi("123")

	fmt.Println(num - 122)

	// string para boolean
	// 1 é verdadeiro, qual quer outro valor e falso
	// True, true, TRUE, 1 são verdadeiros
	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}
}
