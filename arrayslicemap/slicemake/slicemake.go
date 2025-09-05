package main

import "fmt"

func main() {
	s := make([]int, 10) // cria um slice de inteiros com 10 posições
	s[9] = 12            // atribui o valor 12 na posição 9 do slice
	fmt.Println(s)

	s = make([]int, 10, 20) // cria um slice de inteiros com 10 posições e capacidade de 20 posições
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 5, 6, 7, 8, 9, 0) // adiciona os valores 1, 2, 3, 5, 6, 7, 8, 9, 0 ao slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // adiciona o valor 1 ao slice
	fmt.Println(s, len(s), cap(s))
}
