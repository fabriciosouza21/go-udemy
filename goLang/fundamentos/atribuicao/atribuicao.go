package main

import "fmt"

func main() {
	var a byte = 3
	fmt.Println(a)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3 um atalho semelhante em outras linguagens
	i -= 3 // i = i - 3 um atalho semelhante em outras linguagens
	i /= 2 // i = i / 2 um atalho semelhante em outras linguagens
	i *= 2 // i = i * 2 um atalho semelhante em outras linguagens
	i %= 2 // i = i % 2 um atalho semelhante em outras linguagens

	fmt.Println(i)

	// atribuição de múltiplos valores
	x, y := 1, 2
	fmt.Println(x, y)

	// swap
	x, y = y, x

	fmt.Println(x, y)
}
