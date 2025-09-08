package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Println(pessoa, "fala:", texto)
	}
	fmt.Println(pessoa, "terminou")
}

func main() {

	// fale("Maria", "Hello", 3)
	// fale("João", "Goodbye", 3)

	// go fale("Anna", "Olá", 100)
	// go fale("Alef", "Oi", 100)

	go fale("Junior", "Tchau", 100)
	fale("Camila", "Até logo", 100)

}
