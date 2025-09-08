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

func encaminhar(origem <-chan string, destino chan<- string) {
	for {
		destino <- <-origem
	}
}

// multiplexar -  Juntar os channels
func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	// Criando um canal para receber títulos
	channel := make(chan string)
	go encaminhar(entrada1, channel)
	go encaminhar(entrada2, channel)
	return channel
}

func main() {
	// Criando dois canais
	urls := []string{
		"https://www.cod3r.com.br",
		"https://www.alura.com.br",
	}

	// Criando outro canal
	urls2 := []string{
		"https://www.dataactivity.com.br",
		"https://www.google.com",
	}

	// Usando a função juntar para mesclar os channels
	channel := multiplexar(
		title(urls...),
		title(urls2...),
	)

	fmt.Println("Primeiro: ", <-channel)
	fmt.Println("Segundo: ", <-channel)
	fmt.Println("Terceiro: ", <-channel)
	fmt.Println("Quarto: ", <-channel)
}
