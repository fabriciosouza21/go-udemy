package main

import "fmt"

type sportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// método de extensão
// todas as vezes que precisar editar sempre passar o ponteiro
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro := ferrari{"F40", false, 0}
	carro.ligarTurbo()

	// quando o receiver é um ponteiro, a alteração é feita no objeto original
	// por isso para utilizar a interface é necessário passar o ponteiro
	// &ferrari{"F40", false, 0}
	// é mais comun utilizar a interface sem efeito colateral
	// tentar evitar interfaces com efeito colateral
	var carroEsportivo sportivo = &ferrari{"F40", false, 0}
	carroEsportivo.ligarTurbo()

	fmt.Println(carro)
	fmt.Println(carroEsportivo)
}
