package main

import (
	"fmt"
	"time"
)

func main() {

	// A comparação de valores em Go é feita com os operadores relacionais
	// ==, !=, <, >, <=, >=
	// não por referencia de memória

	fmt.Println("Strings: ", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas: ", d1 == d2)
	fmt.Println("Datas: ", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println("Pessoas: ", p1 == p2)

}
