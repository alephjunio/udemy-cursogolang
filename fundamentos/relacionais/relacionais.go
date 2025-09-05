package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("strings:", "Banana" == "Banana") // Comparação de igualdade entre strings
	fmt.Println("!=", 3 != 4)                     // Verifica se os números são diferentes
	fmt.Println(">", 3 > 4)                       // Verifica se o primeiro número é maior
	fmt.Println("<", 3 < 4)                       // Verifica se o primeiro número é menor
	fmt.Println(">=", 3 >= 4)                     // Verifica se é maior ou igual
	fmt.Println("<=", 3 <= 4)                     // Verifica se é menor ou igual

	d1 := time.Unix(0, 0) // Inicializa primeira data
	d2 := time.Unix(0, 0) // Inicializa segunda data

	fmt.Println("Datas:", d1 == d2)     // Compara datas usando operador de igualdade
	fmt.Println("Datas:", d1 != d2)     // Compara datas usando operador de diferença
	fmt.Println("Datas:", d1.Equal(d2)) // Compara datas usando método Equal

}
