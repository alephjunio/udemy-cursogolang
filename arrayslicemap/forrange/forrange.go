package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //compilador que faz a contagem de posições

	// Neste exemplo, usamos o for com range para iterar sobre o array 'numeros'.
	// O range retorna dois valores em cada iteração:
	// - i: o índice atual do elemento no array
	// - numero: o valor do elemento na posição atual
	// Desta forma, podemos acessar tanto a posição quanto o valor de cada elemento do array
	for i, numero := range numeros {
		fmt.Printf("Posição: %d, valor: %d\n", i, numero)
	}

	// Neste caso, usamos o _ (underscore) para ignorar o índice retornado pelo range,
	// já que só queremos iterar sobre os valores do array. O range retorna dois valores:
	// o índice e o valor, mas como só precisamos do valor, descartamos o índice com _
	for _, numero := range numeros {
		fmt.Printf("Valor: %d\n", numero)
	}
}
