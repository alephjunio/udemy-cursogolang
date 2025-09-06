package main

import "fmt"

// Em Go, interfaces são coleções de assinaturas de métodos que um tipo pode implementar.
// A implementação é implícita - se um tipo possui todos os métodos de uma interface,
// ele automaticamente implementa essa interface.

// Interface que define um comportamento de carro luxuoso
// Qualquer tipo que implementar o método fazerBaliza() automaticamente implementa esta interface
type luxuoso interface {
	fazerBaliza()
}

// Interface que define um comportamento de carro esportivo
// Qualquer tipo que implementar o método ligarTurbo() automaticamente implementa esta interface
type esportivo interface {
	ligarTurbo()
}

// Interface composta que combina as interfaces esportivo e luxuoso
// Este é um exemplo de composição de interfaces em Go
// Um tipo que implementa esta interface precisa implementar todos os métodos
// de ambas as interfaces: fazerBaliza() e ligarTurbo()
type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

// Struct que representa um BMW M7
// Esta struct implementará implicitamente todas as interfaces acima
// ao definir os métodos necessários
type bmwM7 struct{}

// Implementação do método fazerBaliza para bmwM7
// Isso faz com que bmwM7 implemente a interface luxuoso
func (b bmwM7) fazerBaliza() {
	fmt.Println("Fazendo baliza")
}

// Implementação do método ligarTurbo para bmwM7
// Isso faz com que bmwM7 implemente a interface esportivo
func (b bmwM7) ligarTurbo() {
	fmt.Println("Ligando turbo")
}

// Como bmwM7 implementa tanto fazerBaliza() quanto ligarTurbo(),
// ela automaticamente implementa a interface esportivoLuxuoso
func main() {
	bmw := bmwM7{}
	bmw.fazerBaliza()
	bmw.ligarTurbo()
}
