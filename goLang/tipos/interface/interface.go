package main

import "fmt"

// interfaces são implementadas implicitamente

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// basta que o tipo implemente o método compatível com a interface
// para que ele seja considerado do tipo da interface

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	// é preciso criar uma variável do tipo da interface
	// para que a função possa ser chamada
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// so funciona pq produto é compatível com a interface imprimivel
	coisa = produto{"Calça Jeans", 79.90}
	imprimir(coisa)

	p2 := produto{"Calça Jeans", 179.90}
	// funciona pq produto é compatível com a interface imprimivel
	imprimir(p2)

}
