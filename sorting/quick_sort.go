package sorting

// Partition selects a pivot, rearranges the array such that all the elements
// less than the pivot are towards the left of it and all the elements greater
// than the pivot are towards the right of it, then returns the index of the
// pivot after rearrangement
func partition(array []int, start, end int) int {
	partitionIndex, pivot := start, array[end]

	for i := start; i < end; i++ {
		if array[i] <= pivot {
			temp := array[i]
			array[i] = array[partitionIndex]
			array[partitionIndex] = temp
			partitionIndex += 1
		}
	}

	// Swap the element at the partition index with the element used as pivot
	temp := array[partitionIndex]
	array[partitionIndex] = array[end]
	array[end] = temp

	return partitionIndex
}

func quickSort(array []int, start, end int) {
	if start < end {
		// Partition the unsorted array
		partitionIndex := partition(array, start, end)

		// Quick sort the segment left to the partition index
		quickSort(array, start, partitionIndex-1)

		// Quick sort the segment right to the partition index
		quickSort(array, partitionIndex+1, end)
	}
}
