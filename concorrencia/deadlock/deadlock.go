package main

import (
	"fmt"
	"time"
)

func rotina(channel chan int) {
	time.Sleep(time.Second)
	channel <- 1 // operação bloqueante
	fmt.Println("Só depois que for lido!!!")
}

func main() {
	channel := make(chan int)

	go rotina(channel)

	fmt.Println(<-channel) // operação bloqueante
	fmt.Println("Foi Lido!")

	fmt.Println(<-channel) // deadlock
	fmt.Println("Fim do Programa.")
}
