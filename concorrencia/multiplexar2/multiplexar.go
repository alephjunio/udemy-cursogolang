package main

import (
	"fmt"
	"time"
)

// falar é uma função que recebe uma pessoa como parâmetro e retorna um canal de saída
// que recebe strings
// A função falar cria um canal de saída, que é retornado para o chamador
// A função falar cria uma goroutine, que é responsável por enviar valores no canal
// A função falar envia valores no canal, a cada segundo, até que o contador atinja 3
// A função falar fecha o canal, indicando que não há mais valores a serem enviados
func falar(pessoa string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			channel <- fmt.Sprintf("%s, falando: %d", pessoa, i)
		}
	}()
	return channel
}

// juntar é uma função que recebe dois canais de entrada e retorna um canal de saída
// que recebe os valores dos canais de entrada
// O select é utilizado para escolher qual canal receberá o valor
// O case é utilizado para verificar se o canal está pronto para receber o valor
// O default é utilizado para verificar se nenhum canal está pronto para receber o valor
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case valor := <-entrada1:
				channel <- valor
			case valor := <-entrada2:
				channel <- valor
			}
		}
	}()
	return channel
}

func main() {
	channel := juntar(falar("Alef"), falar("João"))
	fmt.Println(<-channel, <-channel)
	fmt.Println(<-channel, <-channel)
	fmt.Println(<-channel, <-channel)

}
