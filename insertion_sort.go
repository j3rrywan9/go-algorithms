package main

import (
	"fmt"
)

func insertionSort(array []int) {
	var value, hole int
	for i := 1; i < len(array); i++ {
		value = array[i]
		hole = i
		for hole > 0 && array[hole - 1] > value {
			array[hole] = array[hole - 1]
			hole -= 1
		}
		array[hole] = value
	}
}

func main() {
	a := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", a)
	insertionSort(a)
	fmt.Println("Sorted array: ", a)
}

