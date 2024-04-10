package main

// segundo e primeiro são os nomes dos retornos
// primeiro e segundo são variaveis de escopo
// a ordem dos retornos é a ordem de retorno
func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	x, y := troca(2, 3)
	println(x, y)

	x, y = troca(7, 0)
	println(x, y)

	segundo, primeiro := troca(1, 9)
	println(segundo, primeiro)
}
