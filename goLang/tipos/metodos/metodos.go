package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// quando o receiver é um ponteiro, a alteração é feita no objeto original
// não é necessário utilizar o & para passar o endereço de memória
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"jose", "silva"}
	p1.setNomeCompleto("joao pereira")
	fmt.Println(p1.nomeCompleto())
	// quando o receiver é um ponteiro, a alteração é feita no objeto original
	// não é necessário utilizar o & para passar o endereço de memória
	p1.setNomeCompleto("rafael pereira")
	fmt.Println(p1.nomeCompleto())
}
