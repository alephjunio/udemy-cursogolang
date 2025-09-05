package main

import "fmt"

func imprimirNota(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

}

func main() {
	imprimirNota(6)
	imprimirNota(7)
}
