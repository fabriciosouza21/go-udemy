package main

import "fmt"

// init é uma função especial do Go
// é executada antes da função main
// é utilizada para inicializações
// não é obrigatória
// é executada apenas uma vez
// não aceita parâmetros e não retorna valores
// é possível ter mais de uma função init
// a ordem de execução é a ordem de declaração
func init() {
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Executando a função main")
}
