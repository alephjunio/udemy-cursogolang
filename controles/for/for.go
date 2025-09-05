package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 20; i += 2 {
		fmt.Printf("%d - ", i)
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf(" | %d é Par => ", i)
		} else {
			fmt.Printf(" | %d é ímpar => ", i)
		}
	}

	fmt.Println("For infinito")

	for {
		fmt.Printf(".")
		time.Sleep(time.Second)
	}

}
