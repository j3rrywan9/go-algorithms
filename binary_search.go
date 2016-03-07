package main

import (
	"fmt"
)

func binarySearch(a []int, x int) int {
	low, high := 0, len(a) - 1

	for low <= high {
		mid := (low + high) / 2

		if x == a[mid] {
			return mid
		} else if x < a[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	a := []int{2, 3, 6, 7, 9, 11, 14}
	x := 11
	index := binarySearch(a, x)
	if index == -1 {
		fmt.Printf("Number %d not found.\n", x)
	} else {
		fmt.Printf("Number %d is at index %d.\n", x, index)
	}
}

