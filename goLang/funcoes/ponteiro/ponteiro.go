package main

func inc1(n int) {
	n++
}

// revisão: ponteiro é representado por um asterisco (*)

func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	// funções em go são sempre chamadas por valor
	// uma copia do valor é passada para a função
	// funções não geram efeitos colaterais
	inc1(n) // por valor
	println(n)

	// revisão: & usado para obter o endereço da variável
	inc2(&n) // por referência
	println(n)

	// notas : priorize o uso de valores ao invés de ponteiros
}
