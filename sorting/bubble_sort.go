package sorting

func bubbleSort(array []int) {
	n := len(array)

	for k := 1; k < n; k++ {
		swapped := false

		for i := 0; i < n-1; i++ {
			if array[i] > array[i+1] {
				temp := array[i+1]
				array[i+1] = array[i]
				array[i] = temp
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
