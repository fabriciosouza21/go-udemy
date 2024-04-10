package main

import "fmt"

func main() {
	// é obrigatório inicializar um map
	// é obrigatorio que o ultimo elemento tenha uma vírgula no caso de formatação em várias linhas
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.0

	// não gera erro
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
