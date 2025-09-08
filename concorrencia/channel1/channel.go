package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 1) // Criando um channel de inteiros com buffer de 1

	channel <- 1 // Enviando o valor de 1 ao channel
	<-channel    // Recebendo o valor do channel

	channel <- 2
	fmt.Println("Imprimindo o valor atribuido ao channel", <-channel)
}
