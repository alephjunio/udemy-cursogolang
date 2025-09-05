package main

import "fmt"

// Esta função 'replace' demonstra o uso de retorno nomeado em Go
// Ela recebe dois parâmetros inteiros (old e new) e retorna dois valores também inteiros
// Os valores de retorno são previamente declarados como 'primeiro' e 'segundo'
// A função troca os valores de entrada, colocando 'new' em 'primeiro' e 'old' em 'segundo'
// Por usar retorno nomeado, não é necessário especificar as variáveis no 'return'
func replace(old, new int) (primeiro, segundo int) {
	primeiro = new
	segundo = old
	return
}

func main() {
	param1, param2 := replace(1, 2)
	fmt.Println(param1, param2)

	old, new := replace(2, 1)
	fmt.Println(old, new)
}
