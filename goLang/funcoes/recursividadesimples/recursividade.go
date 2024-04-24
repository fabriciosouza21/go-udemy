package main

import "fmt"

// função recursiva
// função que chama a si mesma
// utiliza o tipo uint para garantir que o número seja positivo
// uma validação forçada para garantir que o número seja positivo
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		fatorialAnterior := fatorial(n - 1)
		return n * fatorialAnterior
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
	//resultado2 := fatorial(-4) - não compila
}
