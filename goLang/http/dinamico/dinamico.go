package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	// Formata a hora atual
	// 02 - dia
	// 01 - mês
	// 2006 - ano
	// 03 - hora
	// 04 - minuto
	// 05 - segundo
	s := time.Now().Format("02/01/2006 03:04:05")

	// Escreve a hora na resposta
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Escutando")
	// primeiro parâmetro é a porta
	// segundo parâmetro é o handler
	log.Fatal(http.ListenAndServe(":3000", nil))
}
