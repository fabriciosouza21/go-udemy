package main

import "fmt"

// defer é uma instrução que adia a execução de uma função até que a função que a contém retorne
// exemplo comum fecha um arquivo após abri-lo
// exemplo comum é liberar recursos após a execução de uma função
func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!") // será executado após a função retornar
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero

}

func main() {
	// valor alto...
	// fim!
	// 5000
	fmt.Println(obterValorAprovado(6000))
	// valor baixo...
	// fim!
	// 3000
	fmt.Println(obterValorAprovado(3000))
}
