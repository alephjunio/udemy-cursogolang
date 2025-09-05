package main

import (
	"fmt"
)

func somar(a int, b int) int {
	return a + b
}

func imprimir(operador string, valor int) {
	fmt.Printf("A %s é: %d\n", operador, valor)
}
