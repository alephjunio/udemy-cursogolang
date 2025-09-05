package main

import "fmt"

func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2
	comprarSorvete := trabalho1 || trabalho2
	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	comprarTv50, comprarTv32, comprarSorvete := compras(true, true)
	fmt.Println("Comprar TV 50\":", comprarTv50)
	fmt.Println("Comprar TV 32\":", comprarTv32)
	fmt.Println("Comprar Sorvete\":", comprarSorvete)
}
