package main

import (
	"fmt"
	"strconv"
)

// LCOJ No. 66
func plusOne(digits []int) []int {
	var num int
	for i := 0; i < len(digits); i++ {
		num = num * 10 + digits[i]	
	}
	num_str := strconv.Itoa(num + 1)
	var res []int
	for _, char := range(num_str) {
		i, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}

func main() {
	mydigits := []int{3, 8, 2, 5}
	fmt.Printf("Input: %v\n", mydigits)
	mydigitsplusone := plusOne(mydigits)
	fmt.Printf("Output: %v\n", mydigitsplusone)
}

