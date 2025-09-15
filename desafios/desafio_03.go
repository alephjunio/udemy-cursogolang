// ### Desafio 2: Cálculo de Mensalidade com Desconto
// **Contexto**: Escolas aplicam descontos progressivos para irmãos.
// **Problema**: Calcule o valor total das mensalidades de uma família considerando:
// - 1º filho: valor integral
// - 2º filho: 10% desconto
// - 3º filho em diante: 20% desconto
// **Input**: `[450.00, 450.00, 380.00, 380.00]` (valores base)
// **Output**: `1598.00` (total com descontos)

package main

import "fmt"

func main() {
	filhos := []float64{450.00, 450.00, 380.00, 380.00}
	total := 0.0
	for i, filho := range filhos {
		if i == 0 {
			total += filho
		} else if i == 1 {
			total += filho * 0.9
		} else {
			total += filho * 0.8
		}
	}
	fmt.Printf("Total com desconto para todos os filhos: R$ %.2f\n", total)
}
