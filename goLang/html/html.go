package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func Titulo(urls ...string) <-chan string {
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
