package main

import (
	"fmt"
)

/*
/ Closure é uma função que retorna outra função
/ A função retornada tem acesso às variáveis do escopo da função pai
/ Isso permite que a função retornada tenha um estado persistente
/ e possa ser usada para criar funções com estado fixo de escopo
*/
func closure() func() {
	x := 10
	var f = func() {
		fmt.Println(x)
	}
	return f
}

func main() {
	x := 20
	fmt.Println(x)

	// closure retorna uma função que imprime o valor de x
	imprimeX := closure()
	imprimeX()
}
