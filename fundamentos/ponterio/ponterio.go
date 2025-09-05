package main

import "fmt"

func main() {
	i := 1 // declaração de variável - declara a variável i do tipo int e atribui o valor 1

	var p *int = new(int) // alocando memória para o ponteiro

	p = &i // pegando o valor da memória da variável i
	*p++   // incrementando o valor da variável i através do ponteiro
	i++    // incrementando o valor da variável i

	fmt.Println("Espaço de Memoria do Ponteiro", p)  // endereço de memória do ponteiro
	fmt.Println("Espaço de Memoria da Variável", &i) // endereço de memória da variável
	fmt.Println("Valor do Ponteiro", *p)             // valor da variável apontada pelo ponteiro
	fmt.Println("Valor da Variável", i)              // valor da variável i

}
