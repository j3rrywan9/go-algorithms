package main

import (
	"fmt"
)

// LC 66
func plusOne(digits []int) []int {
	n := len(digits)

	for index := n - 1; index >= 0; index-- {
		if digits[index] == 9 {
			digits[index] = 0
		} else {
			digits[index]++
			return digits
		}
	}

	digitsPlusOne := make([]int, n+1)
	digitsPlusOne[0] = 1

	return digitsPlusOne
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Printf("Input: %v\n", digits)
	digitsPlusOne := plusOne(digits)
	fmt.Printf("Output: %v\n", digitsPlusOne)

	digits = []int{4, 3, 2, 1}
	fmt.Printf("Input: %v\n", digits)
	digitsPlusOne = plusOne(digits)
	fmt.Printf("Output: %v\n", digitsPlusOne)
}
