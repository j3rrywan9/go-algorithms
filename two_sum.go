package main

import (
	"fmt"
)

func integerInSlice(i int, s []int) bool {
	for _, x := range s {
		if x == i {
			return true
		}
	}
	return false
}

func sliceOfKeys(m map[int]int) []int {
	keys := make([]int, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

// LCOJ No. 1
func twoSum(numbers []int, target int) (int, int) {
	d := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		ks := sliceOfKeys(d)
		if integerInSlice(numbers[i], ks) == false {
			d[target - numbers[i]] = i + 1
		} else {
			return d[numbers[i]], i + 1
		}
	}
	return -1, -1
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Input: %v\n", input)
	fmt.Printf("Target: %d\n", target)
	index1, index2 := twoSum(input, target)
	fmt.Printf("Index1=%d, index2=%d\n", index1, index2)
}

