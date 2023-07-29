package sorting

func selectionSort(array []int) {
	// We only need to run this loop len(array) - 1 times
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		// Swap array[i] and array[min]
		temp := array[i]
		array[i] = array[min]
		array[min] = temp
	}
}
