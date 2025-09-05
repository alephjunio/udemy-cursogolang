package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numberRandom() int {
	s := rand.NewSource(time.Now().UnixNano()) // Cria uma nova fonte de números aleatórios baseada no timestamp atual em nanosegundos
	r := rand.New(s)                           // Cria um novo gerador de números aleatórios usando a fonte criada
	return r.Intn(10)                          // Retorna um número aleatório entre 0 e 9
}

func main() {
	fmt.Println("Número aleatório:", numberRandom()) // Imprime o número aleatório gerado

	if i := numberRandom(); i > 5 { // Se o número aleatório gerado for maior que 5
		fmt.Println("Ganhou!!!", i) // Imprime a mensagem de ganho e o número aleatório gerado
	} else {
		fmt.Println("Perdeu!!!", i) // Imprime a mensagem de perda e o número aleatório gerado
	}
}
