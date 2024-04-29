package main

import (
	"fmt"
	"time"
)

// channel (canal) - é a forma de comunicação entre goroutines
// é um ponto de sincronização
// é um tipo, como qualquer outro tipo em Go

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)

	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)

	c <- 3 * base // enviando dados para o canal

	time.Sleep(3 * time.Second)
	c <- 4 * base // enviando dados para o canal
}

func main() {
	c := make(chan int)

	go doisTresQuatroVezes(2, c)
	fmt.Println("A")
	// recebe os dois primeiros valores do canal, atribuidos no canal
	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	// ler o terceiro valor do canal
	fmt.Println(<-c)

	// error por que não vai ser enviado mais valores para o canal
	// fmt.Println(<-c) // deadlock
}
