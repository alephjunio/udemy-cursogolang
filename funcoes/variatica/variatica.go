package main

import (
	"fmt"
)

// media é uma função que calcula a média de um número variável de argumentos
// numeros é um parâmetro variático do tipo float64
// media retorna a média dos argumentos passados
// Se nenhum argumento for passado, retorna 0
func media(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.1f\n", media(10, 8.5, 9.5))
}
