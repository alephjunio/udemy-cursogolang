package main

import (
	"fmt"
)

/*
// Fatorial é uma operação matemática que consiste na multiplicação de um número natural pelos seus antecessores,
// exceto o zero. Por exemplo, o fatorial de 5 (escrito como 5!) é: 5 x 4 x 3 x 2 x 1 = 120.
// O fatorial de 0 é definido como 1.
// Fatorial calcula o fatorial de um número inteiro
// n é o número inteiro para calcular o fatorial
// retorna o fatorial de n e um possível erro
*/
func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("%d não é possível calcular fatorial de número negativo", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
