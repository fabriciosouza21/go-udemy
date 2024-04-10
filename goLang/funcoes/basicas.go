package main

import "fmt"

// não gerar efeitos colaterais
// função pura não altera o valor de variáveis fora do seu escopo

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println("F2:", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno1", "Retorno2"
}

func main() {
	f1()
	f2("param1", "param2")
	fmt.Println(f3())
	fmt.Println(f4("param1", "param2"))
	r51, r52 := f5()
	_, r53 := f5()
	fmt.Println("F5:", r51, r52)
	fmt.Println("F5:", r53)

}
