package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	//bloqueia até que o buffer tenha espaço para armazenar o valor
	// deadlock
	fmt.Println("Executou")
	ch <- 6
}

func main() {
	// canal com buffer
	// quando o buffer é preenchido, a operação de escrita no canal é bloqueada somente quando o buffer estiver cheio
	// quando o buffer está vazio, a operação é bloqueante
	ch := make(chan int, 3)
	go rotina(ch)
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
