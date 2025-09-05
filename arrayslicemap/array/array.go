package main

import (
	"fmt"
)

func main() {
	var notas [3]float64 // Array com 3 elementos float64, tamanho fixo imutavél
	fmt.Println(notas)

	notas[0] = 10
	notas[1] = 8.5
	notas[2] = 9.5
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas)) // Cálculo da média
	fmt.Printf("Média: %.2f\n", media)   // Imprime a média com 2 casas decimais
}
