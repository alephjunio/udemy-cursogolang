package main

import (
	"fmt"
)

// function1 é uma função sem parâmetro e sem retorno
func function1() {
	fmt.Println("Hello World!")
}

// function2 é uma função com dois parâmetros e sem retorno
func function2(p1 string, p2 string) {
	fmt.Println(p1, p2)
}

// function3 é uma função sem parâmetro e com retorno
func function3() string {
	return "Function 3"
}

// function4 é uma função com dois parâmetros e com retorno
func function4(p1, p2 string) string {
	return fmt.Sprintf("%s %s", p1, p2)
}

// function5 é uma função sem parâmetro e com dois retorno
func function5() (string, string) {
	return "Function 5", "Function 5"
}

// function6 é uma função com dois parâmetros e com dois retorno
func function6(p1, p2 string) (string, string) {
	return p1, p2
}

func main() {
	function1()
	function2("Hello", "World!")
	fmt.Println(function3())
	fmt.Println(function4("Hello", "World!"))
	fmt.Println(function5())
	fmt.Println(function6("Hello", "World!"))

	// r3 e r4 são do tipo string
	r3, r4 := function3(), function4("Hello", "World!")
	fmt.Println("F3 e F4", r3, r4)

	// Retornando dois valores do tipo string.
	// Para usar apenas um dos valores, use _ para ignorar o outro.
	// Sempre é necessario usar ou ignorar um dos valores, colocando de forma explicita.
	// Se não for usar um dos valores, o compilador acusará um erro.
	r5, r6 := function5()
	fmt.Println("F5 e F6", r5, r6)
}
