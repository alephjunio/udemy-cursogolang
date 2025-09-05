package main

import "fmt"

func main() {
	// Criando um map com valores
	funcSalarios := map[string]map[string]float64{
		"SP": {
			"José João":      11325.45,
			"Gabriela Silva": 15456.78,
			"Pedro Junior":   1200.0,
		},
		"RJ": {
			"José João":      11325.45,
			"Gabriela Silva": 15456.78,
			"Pedro Junior":   1200.0,
		},
		"MG": {
			"José João":      11325.45,
			"Gabriela Silva": 15456.78,
			"Pedro Junior":   1200.0,
		},
	}
	// Removendo um elemento do map
	delete(funcSalarios, "RJ")
	fmt.Println(funcSalarios)

	// Imprimindo o map com os valores
	//  Utilizando for range
	for uf, funcs := range funcSalarios {
		fmt.Println(uf, funcs)
	}

	// Imprimindo o map com os valores
	// Utilizando for range ainhado dentro de outro for range
	for uf, funcs := range funcSalarios {
		fmt.Println(uf)
		for nome, salario := range funcs {
			fmt.Printf("\tName: %s, Salary: %.2f\n", nome, salario)
		}
	}
}
