package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	var slice []int

	// array = append(array, 4, 5, 6) -- Não funciona, array é um array com tamanho fixo
	slice = append(slice, 4, 5, 6) // Funciona, slice é um slice com tamanho variável
	fmt.Println(array, slice)

	slice2 := make([]int, 2) // Cria um slice com 2 posições
	copy(slice2, slice)      // Copia os elementos de slice para slice2
	fmt.Println(slice2)
}
