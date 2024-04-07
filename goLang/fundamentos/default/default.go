package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int
	// valores padrões
	// 0 para inteiros
	// 0.0 para float
	// false para boolean
	// "" para string
	// nil para ponteiros

	fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e)
}
