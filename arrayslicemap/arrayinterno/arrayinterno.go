package main

import "fmt"

func main() {
	slice := make([]int, 10, 20)
	slice2 := append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice, slice2)

	slice[0] = 100
	fmt.Println(slice, slice2)
}
