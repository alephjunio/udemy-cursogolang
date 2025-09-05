package main

import "fmt"

func aprovado(nota float64) string {
	if nota >= 7 {
		return "Aprovado"
	}

	return "Reprovado"

}

func main() {
	nota := 7.0
	fmt.Println(aprovado(nota))
}
