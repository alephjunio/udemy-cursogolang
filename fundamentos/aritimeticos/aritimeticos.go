package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma é =>", a+b)          // Soma dois números
	fmt.Println("Subtração é =>", a-b)     // Subtrai um número do outro
	fmt.Println("Multiplicação é =>", a*b) // Multiplica dois números entre si
	fmt.Println("Divisão é =>", a/b)       // Divide o primeiro número pelo segundo

	// bitwise
	fmt.Println("AND =>", a&b)          // Operação AND bit a bit - retorna 1 apenas se ambos bits forem 1
	fmt.Println("OR =>", a|b)           // Operação OR bit a bit - retorna 1 se qualquer bit for 1
	fmt.Println("XOR =>", a^b)          // Operação XOR bit a bit - retorna 1 se os bits forem diferentes
	fmt.Println("NOT =>", ^a)           // Inverte todos os bits do número
	fmt.Println("Left Shift =>", a<<b)  // Desloca bits para esquerda - multiplica por 2^b
	fmt.Println("Right Shift =>", a>>b) // Desloca bits para direita - divide por 2^b

	// Outras operações usando a biblioteca math
	c := 3.0
	d := 2.0

	fmt.Println("Potência =>", math.Pow(c, d))                 // Eleva o primeiro número pelo segundo
	fmt.Println("Radical =>", math.Sqrt(c))                    // Calcula a raiz quadrada do número
	fmt.Println("Arredondamento =>", math.Round(c))            // Arredonda para o inteiro mais próximo
	fmt.Println("Arredondamento para cima =>", math.Ceil(c))   // Arredonda sempre para cima
	fmt.Println("Arredondamento para baixo =>", math.Floor(c)) // Arredonda sempre para baixo
	fmt.Println("Maior =>", math.Max(c, d))                    // Retorna o maior entre dois números
	fmt.Println("Menor =>", math.Min(c, d))                    // Retorna o menor entre dois números
	fmt.Println("Resto de Divisão =>", math.Mod(c, d))         // Retorna o resto da divisão entre dois números

}
