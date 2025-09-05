package main

import (
	"fmt"
)

/*
A função aprovados é uma função variática que recebe um número variável de argumentos do tipo string.
A função imprime a lista de aprovados, com o número do aluno e o nome.
Se não houver aprovados, imprime uma mensagem informando.
Se houver apenas um aprovado, imprime uma mensagem informando.
Se houver mais de um aprovado, imprime uma mensagem informando.
*/
func aprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}

}

func main() {
	listaAprovados := []string{"Alef", "Bruno", "Carlos"}
	listaAprovados = append(listaAprovados, "Daniel")
	listaAprovados = append(listaAprovados, "Eduardo")
	listaAprovados = append(listaAprovados, "Fábio")

	aprovados(listaAprovados...)

}
