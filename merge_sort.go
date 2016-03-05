package main

import (
	"fmt"
)

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	left, right := []int{}, []int{}

	mid := len(array) / 2

	for i, x := range array {
		switch {
		case i < mid:
			left = append(left, x)
		case i >= mid:
			right = append(right, x)
		}
	}

	left, right = mergeSort(left), mergeSort(right)

	return merge(left, right)	
}

// Merge two sorted arrays into one sorted array
func merge(left, right []int) []int {
	array := make([]int, len(left) + len(right))

	i, j, k := 0, 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			array[k] = left[i]
			i += 1
		} else {
			array[k] = right[j]
			j += 1
		}
		k += 1
	}

	for i < len(left) {
		array[k] = left[i]
		i += 1
		k += 1
	}

	for j < len(right) {
		array[k] = right[j]
		j += 1
		k += 1
	}

	return array
}

func main() {
	a := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", a)
	a = mergeSort(a)
	fmt.Println("Sorted array: ", a)
}

