package main

import (
	"fmt"
)

func bubbleSort(array []int) {
	for k := 1; k < len(array); k++ {
		flag := 0
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				temp := array[i+1]
				array[i+1] = array[i]
				array[i] = temp
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	}
}

func main() {
	a := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("Unsorted array: ", a)
	bubbleSort(a)
	fmt.Println("Sorted array: ", a)
}
