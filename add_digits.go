package main

import (
	"fmt"
)

// LeetCode OJ NO. 258
func addDigits(num int) int {
	if num == 0 {
		return 0
	} else if num % 9 == 0 {
		return 9
	}
	return num % 9
}

func main() {
	result := addDigits(36)
	fmt.Printf("Input: %d\n", 36)
	fmt.Printf("Result: %d\n", result)
	result = addDigits(0)
	fmt.Printf("Input: %d\n", 0)
	fmt.Printf("Result: %d\n", result)
	result = addDigits(38)
	fmt.Printf("Input: %d\n", 38)
	fmt.Printf("Result: %d\n", result)
}
