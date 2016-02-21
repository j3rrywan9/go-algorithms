package main

import (
	"fmt"
	"strconv"
)

// LeetCode OJ No. 66
func plusOne(digits []int) int {
	var num int

	// Get the number represented by the digits
	for i := 0; i < len(digits); i++ {
		num = num * 10 + digits[i]	
	}

	// Convert (number + 1) to string
	num_str := strconv.Itoa(num + 1)

	var digits_plus_one []int

	for _, char := range(num_str) {
		// Convert each character in the string to integer
		i, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		digits_plus_one = append(digits_plus_one, i)
	}

	var num_plus_one int

	for i := 0; i < len(digits_plus_one); i++ {
		num_plus_one = num_plus_one * 10 + digits_plus_one[i]	
	}

	return num_plus_one 
}

func main() {
	digits := []int{3, 8, 2, 5}
	fmt.Printf("Input: %v\n", digits)
	digits_plus_one := plusOne(digits)
	fmt.Printf("Output: %d\n", digits_plus_one)
}

