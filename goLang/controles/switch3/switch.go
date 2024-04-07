package main

import "time"

// tipo retorna o tipo da variável
// interface{} é um tipo que aceita qualquer tipo
// switch com type
// i.(type) é um assert que retorna o tipo da variável

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main() {
	println(tipo(2))
	println(tipo(2.3))
	println(tipo("Opa"))
	println(tipo(func() {}))
	println(tipo(time.Now()))
}
