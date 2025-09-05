package main

import (
	"fmt"
	"time"
)

func main() {

	time := time.Now()
	hour := time.Hour()

	switch {
	case hour < 12:
		fmt.Println("Bom dia!")
	case hour < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}

}
