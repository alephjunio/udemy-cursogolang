package main

import (
	"fmt"
	"time"
)

func rotina(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	channel <- 5
	fmt.Println("Enviado 5")
	channel <- 6
	channel <- 7
}

func main() {
	channel := make(chan int, 3)

	go rotina(channel)

	time.Sleep(time.Second)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
