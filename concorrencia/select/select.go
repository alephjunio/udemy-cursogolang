package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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

func oMaisRapido(url1, url2, url3 string) string {
	channel1 := title(url1)
	channel2 := title(url2)
	channel3 := title(url3)

	// Estrutura de controle de fluxo de concorrencia
	select {
	case t1 := <-channel1:
		return t1
	case t2 := <-channel2:
		return t2
	case t3 := <-channel3:
		return t3
	case <-time.After(1 * time.Second):
		return "Nenhum dos sites respondeu em 1 segundo"
	}
}

func main() {

	vencedor := oMaisRapido(
		"https://www.cod3r.com.br",
		"https://www.alura.com.br",
		"https://www.dataactivity.com.br",
	)

	fmt.Println(vencedor)

}
