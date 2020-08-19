package sorting

func insertionSort(array []int) {
	var value, hole int

	for i := 1; i < len(array); i++ {
		value = array[i]
		hole = i

		for hole > 0 && array[hole-1] > value {
			array[hole] = array[hole-1]
			hole -= 1
		}

		array[hole] = value
	}
}
