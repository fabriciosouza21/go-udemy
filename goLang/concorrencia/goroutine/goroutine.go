package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Por que você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// go criar uma goroutine
	// thread leve
	// thread são linhas de execução independentes
	// goroutine não são threads
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)

	// quando a thread principal termina, todas as goroutines são encerradas

	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!", 5)
	// fmt.Println("Fim!")
}
