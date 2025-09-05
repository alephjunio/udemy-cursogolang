package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 7 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println("Nota 10:", notaParaConceito(10))
	fmt.Println("Nota 7:", notaParaConceito(7))
	fmt.Println("Nota 5:", notaParaConceito(5))
	fmt.Println("Nota 3:", notaParaConceito(3))
	fmt.Println("Nota 1:", notaParaConceito(1))
}
