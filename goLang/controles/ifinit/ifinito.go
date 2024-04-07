package main

import (
	"math/rand"
	"time"
)

// if init; condition { }
func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	// bloco de inicialização
	// o valor do i só é válido dentro do if e else
	// o valor fica disponivel somente no contexto do bloco
	if i := numeroAleatorio(); i > 5 {
		println("Ganhou!!!")
	} else {
		println("Perdeu!!!")
	}

	// println(i) // i não está mais disponível
}
