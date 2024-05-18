package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// google I/O 2012 - Go Concurrency Patterns

// <- chan - canal somente leitura
// generator retorna um canal somente leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		// função anônima
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[0]
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.udemy.com/", "https://www.youtube.com")

	// forma de receber dados de canais

	fmt.Println("Primeiros:", <-t1, " | ", <-t2)
	fmt.Println("Segundos:", <-t1, " | ", <-t2)
}
