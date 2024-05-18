package main

import (
	"fmt"

	"curso.com/m/html"
)

// multiplexar é a técnica de enviar várias mensagens em um único canal.
// <-chan - canal somente leitura
// quando não tem a setar o canal é bidirecional
func encaminhar(origem <-chan string, destiono chan string) {
	for {
		destiono <- <-origem
	}
}

// multiplexar - misturar (mensagens) em um único canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.udemy.com/", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
