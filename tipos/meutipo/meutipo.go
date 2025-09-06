package main

import "fmt"

// notaParaConceito retorna o conceito correspondente à nota passada como parâmetro
// nota é um tipo alias para float64
// nota.entre é um método que verifica se a nota está entre dois valores
// notaParaConceito é uma função que retorna o conceito correspondente à nota passada como parâmetro
type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func notaParaConceitoSwitch(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 6.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println("Nota 10:", notaParaConceito(10.0))
	fmt.Println("Nota 7:", notaParaConceito(7.5))
	fmt.Println("Nota 5:", notaParaConceitoSwitch(5.0))
	fmt.Println("Nota 3:", notaParaConceitoSwitch(3.2))
	fmt.Println("Nota 1:", notaParaConceitoSwitch(1.3))
}
