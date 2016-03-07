package main

import (
	"fmt"
)

func binarySearch(array []int, x int) int {
	low, high := 0, len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		switch {
		case x == array[mid]:
			return mid
		case x < array[mid]:
			high = mid - 1
		case x > array[mid]:
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

