package main

import "fmt"

func notaParaConceito(nota float64) string {
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println("Nota:", notaParaConceito(10))
	fmt.Println("Nota:", notaParaConceito(9))
	fmt.Println("Nota:", notaParaConceito(8))
	fmt.Println("Nota:", notaParaConceito(7))
	fmt.Println("Nota:", notaParaConceito(6))
	fmt.Println("Nota:", notaParaConceito(5))
	fmt.Println("Nota:", notaParaConceito(4))
	fmt.Println("Nota:", notaParaConceito(3))
	fmt.Println("Nota:", notaParaConceito(2))
	fmt.Println("Nota:", notaParaConceito(1))
	fmt.Println("Nota:", notaParaConceito(-1))
	fmt.Println("Nota:", notaParaConceito(11))
}
