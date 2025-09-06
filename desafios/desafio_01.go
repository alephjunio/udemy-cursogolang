package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numIndexMap := make(map[int]int)
	for i, num := range nums {
		numIndexMap[num] = i
	}

	for i, num := range nums {
		complement := target - num

		if index, found := numIndexMap[complement]; found && index != i {
			return []int{i, index}
		}
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
