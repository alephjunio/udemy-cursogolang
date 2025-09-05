package main

import "fmt"

func main() {
	// Maps precisam ser inicializados antes de serem utilizados
	approved := make(map[int]string) // Inicialização de um map

	approved[12345678] = "Maria"
	approved[98765432] = "Pedro"
	approved[95135745] = "Ana"

	fmt.Println(approved) // Imprimindo o map

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approved[12345678])
	delete(approved, 12345678)                                   // Deletando um elemento do map
	fmt.Println("Removido: Chave 12345678 ", approved[12345678]) // Imprimindo o map após a exclusão
}
