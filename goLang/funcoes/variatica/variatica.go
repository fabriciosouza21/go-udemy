package main

import "fmt"

// variadic function
// é a forma de passar um número variável de parâmetros para uma função
// o operador ... é usado para indicar que a função aceita um número variável de parâmetros

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 10.0))

}
