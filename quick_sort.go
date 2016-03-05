package main

import (
	"fmt"
)

// Partition selects a pivot, rearranges the array such that all the elements
// less than the pivot are towards the left of it and all the elements greater
// than the pivot are towards the right of it, then returns the index of the
// pivot after rearrangement
func partition(array []int, start, end int) int {
	index_of_pivot, pivot := start, array[end]

	for i := start; i < end; i++ {
		if array[i] <= pivot {
			temp := array[i]
			array[i] = array[index_of_pivot]
			array[index_of_pivot] = temp
			index_of_pivot += 1
		}
	}
	temp := array[index_of_pivot]
	array[index_of_pivot] = array[end]
	array[end] = temp

	return index_of_pivot
}

func quickSort(array []int, start, end int) {
	if start < end {
		// Partition the unsorted array
		index_of_pivot := partition(array, start, end)

		// Quick sort the segment left to the pivot index
		quickSort(array, start, index_of_pivot - 1)

		// Quick sort the segment right to the pivot index
		quickSort(array, index_of_pivot + 1, end)
	}
}

func main() {
	a := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", a)
	quickSort(a, 0, len(a) - 1)
	fmt.Println("Sorted array: ", a)
}

