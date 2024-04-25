package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// método: função com receiver (receptor)
// receiver é um parâmetro de um tipo
// função associada a um tipo
// método é uma função
// método é uma função com receiver
// receiver é um parâmetro
// receiver é um parâmetro entre o func e o nome do método
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1939.34, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}
