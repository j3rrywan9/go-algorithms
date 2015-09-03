package main

import (
	"fmt"
)

func partition(a []int, start, end int) int {
	var pIndex int = start
	var pivot int = a[end]
	for i := start; i < end; i++ {
		if a[i] <= pivot {
			temp := a[i]
			a[i] = a[pIndex]
			a[pIndex] = temp
			pIndex += 1
		}
	}
	temp := a[pIndex]
	a[pIndex] = a[end]
	a[end] = temp
	return pIndex
}

func quickSort(a []int, start, end int) {
	if start < end {
		pIndex := partition(a, start, end)
		quickSort(a, start, pIndex - 1)
		quickSort(a, pIndex + 1, end)
	}
}

func main() {
	a := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", a)
	quickSort(a, 0, len(a) - 1)
	fmt.Println("Sorted array: ", a)
}

