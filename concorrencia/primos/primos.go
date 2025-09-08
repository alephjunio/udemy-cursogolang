package main

import (
	"fmt"
	"time"
)

// Função para verificar se um número é primo
func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Função para gerar primos e enviá-los para um canal
func primos(n int, channel chan int) {
	start := 2
	for i := 0; i < n; i++ {
		for primo := start; ; primo++ {
			if isPrimo(primo) {
				channel <- primo
				start = primo + 1
				time.Sleep(time.Millisecond * 200)
			}
		}
	}
	close(channel)
}

func main() {
	// Criando um canal para receber primos
	channel := make(chan int, 30)

	// Iniciando a goroutine para gerar primos
	go primos(cap(channel), channel)

	// Recebendo primos do canal e imprimindo
	for primo := range channel {
		fmt.Printf("%d,", primo)
	}
}
