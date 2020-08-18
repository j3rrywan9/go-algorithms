package search

func binarySearch(array []int, x int) int {
	low, high := 0, len(array)-1

	for low <= high {
		mid := low + (high-low)/2

		switch {
		case x == array[mid]:
			return mid
		case x < array[mid]:
			high = mid - 1
		case x > array[mid]:
			low = mid + 1
		}
	}

	return -1
}
