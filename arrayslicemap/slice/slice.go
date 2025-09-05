package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3} // array é um Array
	slice := []int{1, 2, 3}  // slice apenas ponteiro para uma fatia do array
	fmt.Print("Array:", array, " | ")
	fmt.Print("Slice:", slice, "\n")
	fmt.Print("Tamanho do array:", len(array), " | ")   // len() retorna o tamanho do array
	fmt.Print("Tamanho do slice:", len(slice), " | ")   // len() retorna o tamanho do slice
	fmt.Print("Capacidade do slice:", cap(slice), "\n") // cap() retorna a capacidade do slice

	array2 := [5]int{1, 2, 3, 4, 5}

	/*
		slice não é um Array
	*/
	slice2 := array2[1:3] // slice2 é uma fatia do array2
	fmt.Print("Array2:", array2, " | ")
	fmt.Print("Slice2:", slice2, "\n")

	slice3 := array2[:2] // slice3 é uma fatia do array2
	fmt.Print("Slice3:", slice3, "\n")

	slice4 := slice3[:2] // slice4 é uma fatia do slice3
	fmt.Print("Slice4:", slice4, "\n")
}
