
 ``` go
 package main

  

import "fmt"

  

func main() {

// mesma linha, no final dela não cria \n

fmt.Print("Mesma")

fmt.Print(" linha.")

  

// adicionar uma quebra de linha

fmt.Println(" Nova")

fmt.Println("linha")

  

x := 3.141516

// fmt.Println("O valor de x é " + x) // não compila

  

// fmt.Println("O valor de x é " + fmt.Sprint(x)) // compila

xs := fmt.Sprint(x)

fmt.Println("O valor de x é " + xs)

  

// fmt.Println("O valor de x é", x) // compila

fmt.Println("O valor de x é", x)

  

// fmt.printf permite adicionar ganchos para formatar a string

fmt.Printf("O valor de x é %.2f.", x)

  

a := 1

b := 1.9999

c := false

d := "opa"

  

// %d = int, %f = float, %.1f = float com uma casa decimal, %t = bool, %s = string

  

fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

  

//%v tipos genéricos

fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
 ```