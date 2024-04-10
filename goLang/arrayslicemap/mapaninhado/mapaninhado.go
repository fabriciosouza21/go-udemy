package main

import "fmt"

func main() {
	funcsPorLertra := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 15456.78,
			"Guga Pereira":  8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLertra, "P")

	for let, funcs := range funcsPorLertra {
		fmt.Println(let, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
