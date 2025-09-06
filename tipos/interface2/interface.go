// Em Go, interfaces são conjuntos de métodos que definem um comportamento.
// Uma interface define um contrato que os tipos devem seguir.
// Qualquer tipo que implementa todos os métodos de uma interface, automaticamente satisfaz essa interface.

// Interfaces em Go são implementadas implicitamente - não é necessário declarar explicitamente
// que um tipo implementa uma interface, basta implementar seus métodos.

package main

import "fmt"

// Interface que define o comportamento de um carro esportivo
// Neste caso, apenas um método é necessário: ligarTurbo()
type sportivo interface {
	ligarTurbo()
}

// Struct ferrari que implementará a interface sportivo
// Note que não precisamos declarar explicitamente que ferrari implementa sportivo
type ferrari struct {
	modelo      string
	turboLigado bool
	velocidade  int
}

// Método que implementa a interface sportivo
// Ao implementar este método, ferrari automaticamente satisfaz a interface
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	// Criando uma instância de ferrari diretamente
	f := &ferrari{
		modelo:      "F40",
		turboLigado: false,
		velocidade:  0,
	}

	f.ligarTurbo()

	// Criando uma instância de ferrari através da interface
	// Isso demonstra o polimorfismo em Go - podemos usar qualquer tipo
	// que implemente a interface sportivo
	var f2 sportivo = &ferrari{
		modelo:      "F40",
		turboLigado: false,
		velocidade:  0,
	}
	f2.ligarTurbo()

	fmt.Println(f, f2)
}
