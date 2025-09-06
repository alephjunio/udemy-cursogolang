package main

import (
	"fmt"
)

/*
// O defer é uma palavra chave que permite que atrase a chamada de uma função
// Sendo util para executar uma função após a conclusão de outra função
// O defer é útil para liberar recursos, como fechar arquivos, conexões de banco de dados, etc.
*/

func retornarValor(numero int) int {
	defer fmt.Println("Timer DEFER")
	if numero > 10 {
		return 10
	}

	fmt.Println("Retornando valor abaixo de 10")
	return numero
}

func main() {
	resultado := retornarValor(5)
	fmt.Println(resultado)
}
