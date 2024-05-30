package main

import (
	"log"
	"net/http"

	"curso.com/m/http/servedb/cliente"
)

func main() {
	http.HandleFunc("/usuarios/", cliente.UsuarioHandler)
	log.Println("Escutando")
	http.ListenAndServe(":3000", nil)

}
