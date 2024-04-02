package main

import (
	"fmt"
	"math"
)

func main() {

	// lebrando sempre que go é fortimente tipado, então não é possível fazer operações com tipos diferentes
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)

	fmt.Println("Subtração =>", a-b)

	fmt.Println("Divisão =>", a/b)

	fmt.Println("Multiplicação =>", a*b)

	fmt.Println("Módulo =>", a%b)

	// bitwise

	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operações usando math...
	// é necessário converter para float64

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(float64(a), float64(b)))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}
