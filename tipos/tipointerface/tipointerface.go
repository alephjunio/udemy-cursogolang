package main

import "fmt"

// Interface em Go é um tipo abstrato que define um conjunto de métodos
// Uma interface vazia (interface{}) não possui métodos definidos, portanto pode
// armazenar valores de qualquer tipo

// Exemplo de struct para demonstração
type curso struct {
	nome string
}

func main() {
	// interface{} é o tipo mais genérico em Go
	// Pode armazenar qualquer tipo de valor (int, string, struct, etc)
	var coisa interface{}
	fmt.Println(coisa) // Imprime nil (valor zero de interface)

	// Podemos atribuir qualquer valor à interface{}
	coisa = 3 // Atribuindo um inteiro
	fmt.Println(coisa)

	// Podemos criar um alias para interface{} para melhor legibilidade
	type dinamico interface{}

	// Inicializando com uma string
	var coisa2 dinamico = "Opa!"
	fmt.Println(coisa2)

	// Podemos mudar o tipo do valor armazenado a qualquer momento
	coisa2 = true // Agora armazena um booleano
	fmt.Println(coisa2)

	// Também pode armazenar structs
	coisa2 = curso{
		nome: "Explorando a Lingaguem Go",
	}

	fmt.Println(coisa2)
	// %T mostra o tipo real do valor armazenado na interface
	fmt.Printf("%T\n", coisa2)
}
