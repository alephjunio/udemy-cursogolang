package main

import "fmt"

// qual é a diferença entre uma função e uma função anônima?
// uma função é um bloco de código que pode ser reutilizado
// uma função anônima é uma função sem nome, que pode ser atribuída a uma variável
// uma função anônima pode ser usada como argumento para outra função

var sum = func(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(sum(1, 2))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(1, 2))

	// uma função anônima pode ser usada como valor de retorno
	// nesse caso, a função anônima é retornada pela função 'operacao'
	// e pode ser chamada posteriormente
	operacao := func(a, b int) func(int, int) int {
		return func(a, b int) int {
			return a + b
		}
	}
	fmt.Println(operacao(1, 2)(1, 2))

}
