package main

import (
	"fmt"
)

// Determine if an integer is in a slice of integers
func integerInSlice(i int, s []int) bool {
	for _, x := range s {
		if x == i {
			return true
		}
	}
	return false
}

// Construct a slice of integers that come from keys of a map[int]int
func sliceOfKeys(m map[int]int) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// LeetCode OJ No. 1
func twoSum(nums []int, target int) (int, int) {
	d := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		keys := sliceOfKeys(d)
		if integerInSlice(nums[i], keys) == false {
			d[target-nums[i]] = i
		} else {
			return d[nums[i]], i
		}
	}
	return -1, -1
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Numbers: %v\n", nums)
	fmt.Printf("Target: %d\n", target)
	index1, index2 := twoSum(nums, target)
	fmt.Printf("Index1=%d, index2=%d\n", index1, index2)
}
