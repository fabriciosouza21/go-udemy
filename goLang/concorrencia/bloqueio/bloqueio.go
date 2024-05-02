package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Só depois que foi lido")
}

func main() {
	// quando criar um canal sem buffer, a operação de escrita no canal é bloqueante
	c := make(chan int) // canal sem buffer
	go rotina(c)
	// quadno não tem buffer, a operação de leitura é bloqueante
	//
	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}
