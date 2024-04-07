package main

func imprimirResultado(nota float64) {
	// não é necessário colocar parênteses em torno da condição
	// as chaves são obrigatórias

	if nota >= 7 {
		println("Aprovado com nota", nota)
	} else {
		println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
