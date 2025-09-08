package main

import (
	"fmt"
	"time"
)

// doisTresQuatro é uma função que recebe um valor base e um channel como parâmetros
// e envia os valores 2*base, 3*base e 4*base para o channel com um atraso de 1 segundo entre cada envio
func doisTresQuatro(base int, channel chan int) {
	time.Sleep(time.Second)
	channel <- 2 * base

	time.Sleep(time.Second)
	channel <- 3 * base

	time.Sleep(3 * time.Second)
	channel <- 4 * base
}

func main() {
	channel := make(chan int)     // Criando um channel de inteiros
	go doisTresQuatro(2, channel) // Chamando a função doisTresQuatro em uma goroutine

	a, b := <-channel, <-channel // Recebendo os valores do channel
	fmt.Println(a, b)

	c := <-channel // Recebendo o último valor do channel
	fmt.Println(c)

}
