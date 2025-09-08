package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func title(urls ...string) <-chan string {
	// Criando um canal para receber títulos
	channel := make(chan string)
	// Iniciando uma goroutine para cada URL
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)             // Realizando a chamada HTTP
			html, _ := ioutil.ReadAll(resp.Body) // Lendo o corpo da resposta

			r, _ := regexp.Compile(`<title>(.*?)</title>`) // Compilando a expressão regular
			title := r.FindStringSubmatch(string(html))[1] // Encontrando o título

			channel <- title // Enviando o título para o canal
		}(url)
	}
	return channel
}

func main() {

	urls := title("https://www.cod3r.com.br", "https://www.dataactivity.com.br", "https://www.golang.org")

	fmt.Println("Primeiro:", <-urls)
	fmt.Println("Segundo:", <-urls)
	fmt.Println("Terceiro:", <-urls)
}
