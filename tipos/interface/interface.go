package main

import (
	"fmt"
)

// Em Go, interfaces são coleções de assinaturas de métodos que um tipo pode implementar.
// Uma interface define um contrato de comportamento que os tipos devem seguir.
// A implementação de interfaces em Go é implícita - não é necessário declarar explicitamente
// que um tipo implementa uma interface.

// Interface que define o comportamento de algo que pode ser convertido para string
type imprimivel interface {
	toString() string
}

// Structs são tipos que agrupam campos relacionados
type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// Quando implementamos o método toString() para pessoa,
// automaticamente o tipo pessoa passa a satisfazer a interface imprimivel
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

// O mesmo acontece com produto ao implementar toString()
// Agora produto também satisfaz a interface imprimivel
func (p produto) toString() string {
	return fmt.Sprintf("Produto: %s, Preço: %.2f", p.nome, p.preco)
}

// Esta função aceita qualquer tipo que implemente a interface imprimivel
// Demonstrando o polimorfismo em Go - diferentes tipos podem ser tratados de forma uniforme
func imprimir(param imprimivel) {
	fmt.Println(param.toString())
}

func main() {
	// Podemos atribuir uma pessoa a uma variável do tipo interface
	// já que pessoa implementa todos os métodos da interface
	var coisa imprimivel = pessoa{
		nome:      "Alef",
		sobrenome: "Silva",
	}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// A mesma variável pode receber um produto
	// demonstrando a flexibilidade das interfaces
	coisa = produto{
		nome:  "Notebook",
		preco: 2500.0,
	}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// Podemos passar diretamente uma pessoa para a função
	// sem precisar declarar explicitamente a interface
	p2 := pessoa{
		nome:      "João",
		sobrenome: "Silva",
	}
	imprimir(p2)
}
