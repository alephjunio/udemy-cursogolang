package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

// exec é uma função que executa outra função passada como parâmetro
// f é a função que será executada
// a e b são os parâmetros da função f
// exec retorna o resultado da execução da função f

func exec(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func main() {
	result := exec(multiplicacao, 2, 3)
	fmt.Println(result)
}
