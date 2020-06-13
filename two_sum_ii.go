package main

import (
	"fmt"
)

// LC 167
func twoSumII(numbers []int, target int) []int {
	var l int
	r := len(numbers) - 1

	for numbers[l]+numbers[r] != target {
		if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return []int{l + 1, r + 1}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	result := twoSumII(numbers, 9)
	fmt.Printf("[%d, %d]\n", result[0], result[1])
}
