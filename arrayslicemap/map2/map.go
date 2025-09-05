package main

import "fmt"

func main() {
	// Criando um map com valores
	funcsESalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	// Imprimendo o map com os valores
	fmt.Println(funcsESalarios)

	// Removendo um elemento do map
	delete(funcsESalarios, "Inexistent") // Deletando um elemento inexistente do map
	fmt.Println(funcsESalarios)          // Imprimindo o map após a exclusão

	// Update a value
	funcsESalarios["Gabriela Silva"] = 16500.25

	// Imprimindo o map após a atualização, utilizando o for range
	for nome, salario := range funcsESalarios {
		fmt.Printf("Name: %s, Salary: %.2f\n", nome, salario)
	}
}
