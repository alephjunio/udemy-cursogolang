package main

import "fmt"

// Input: nums = [2, 7, 11, 15], target = 9
// Output: [0, 1] (porque nums[0] + nums[1] = 2 + 7 = 9)

func soma(nums []int, target int) []int {
	fmt.Println("\nIniciando função soma com:")
	fmt.Printf("nums=%v, target=%d\n", nums, target)

	numMap := make(map[int]int)
	fmt.Println("\nMapa vazio inicializado:", numMap)

	for i, num := range nums {
		fmt.Printf("\nIteração %d: Processando num=%d\n", i+1, num)

		complement := target - num
		fmt.Printf("Complemento calculado: %d - %d = %d\n", target, num, complement)

		if index, exists := numMap[complement]; exists {
			fmt.Printf("Encontrado complemento %d no índice %d\n", complement, index)
			fmt.Printf("Retornando par: [%d, %d]\n", index, i)
			return []int{index, i}
		}

		numMap[num] = i
		fmt.Printf("Adicionado ao mapa: numMap[%d] = %d\n", num, i)
		fmt.Println("Estado atual do mapa:", numMap)
	}

	fmt.Println("\nNenhuma solução encontrada, retornando slice vazio")
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println("Iniciando programa com:")
	fmt.Printf("nums=%v\ntarget=%d\n", nums, target)

	result := soma(nums, target)
	fmt.Printf("\nResultado final: nums=%v, target=%d -> %v\n", nums, target, result)
}
