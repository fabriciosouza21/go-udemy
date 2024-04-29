package main

import "fmt"

func main() {
	// criar um canal, um tipo e um buffer
	ch := make(chan int, 1)

	// enviar dados para o canal(equivalente a um processo de escrita)
	ch <- 1

	// receber dados do canal(equivalente a um processo de leitura)
	<-ch

	ch <- 2
	// o canal é m mecanismo de sincronização
	fmt.Println(<-ch)
}
