package main

func notaParaConceito(n float64) string {
	var nota = int(n)

	// no go não é necessário colocar break
	// o case já é um bloco
	// o default é opcional
	// o switch não executa o próximo caso automaticamente
	// para isso é necessário usar a palavra reservada fallthrough
	// o switch só aceita valores inteiros, strings ou expressões que resultem em valores inteiros
	switch nota {
	case 10:
		fallthrough // executa o próximo caso
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	println(notaParaConceito(9.8))
	println(notaParaConceito(6.9))
	println(notaParaConceito(2.1))
}
