package main

import "fmt"

// Interfaces em Go são uma forma de definir comportamentos
// Uma interface é um conjunto de métodos que um tipo deve implementar
//
// Exemplo de interface:
// type Forma interface {
//     Area() float64
//    Perimetro() float64
// }
//
// Para implementar uma interface, um tipo precisa implementar todos os métodos
// definidos nela. Em Go, isso é feito implicitamente - não é necessário
// declarar explicitamente que um tipo implementa uma interface
//
// Interfaces permitem:
// - Polimorfismo
// - Desacoplamento entre componentes
// - Testes mais fáceis através de mocks
// - Maior flexibilidade no código

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}
	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
