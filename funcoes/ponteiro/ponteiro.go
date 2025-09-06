package main

import (
	"fmt"
)

// inc1 é uma função que incrementa o valor de um número inteiro
// Porém, o valor incrementado não é retornado
func inc1(n int) {
	n++
}

// inc2 é uma função que incrementa o valor de um número inteiro
// Porém, o valor incrementado é retornado
func inc2(n *int) {
	// n é um ponteiro para um inteiro
	// *n é o valor apontado por n
	*n++
}

func main() {
	n := 10
	inc1(n)
	fmt.Println(n)
	inc2(&n)
	fmt.Println(n)
}
