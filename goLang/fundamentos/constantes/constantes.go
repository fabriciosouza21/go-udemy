package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador
	area := PI * raio * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 1
		d = 2
	)

	fmt.Println(a, b, c, d)

	g, h, i := 1, false, "opa"
	fmt.Println(g, h, i)

}
