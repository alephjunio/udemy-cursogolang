package main

import (
	"fmt"
)

// Pseudo herança
// Em Go, não existe herança como em outras linguagens orientadas a objetos.

type carro struct {
	nome       string
	velocidade int
}

type esportivo struct {
	carro
	turbo bool
}

func main() {
	c := esportivo{}
	c.nome = "Fusca"
	c.velocidade = 100
	c.turbo = true

	fmt.Println("Nome:", c.nome)
	fmt.Println("Velocidade:", c.velocidade)
	fmt.Println("Turbo:", c.turbo)
}
