package main

import (
	"fmt"
)

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	left := make([]int, 0)
	right := make([]int, 0)

	m := len(a) / 2

	for i, x := range a {
		switch {
		case i < m:
			left = append(left, x)
		case i >= m:
			right = append(right, x)
		}
	}

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)	
}

// Merge two sorted arrays into one
func merge(left, right []int) []int {
	res := make([]int, (len(left) + len(right)))

	i, j, k := 0, 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res[k] = left[i]
			i += 1
		} else {
			res[k] = right[j]
			j += 1
		}
		k += 1
	}
	for i < len(left) {
		res[k] = left[i]
		i += 1
		k += 1
	}
	for j < len(right) {
		res[k] = right[j]
		j += 1
		k += 1
	}
	return res
}

func main() {
	a := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", a)
	a = mergeSort(a)
	fmt.Println("Sorted array: ", a)
}

