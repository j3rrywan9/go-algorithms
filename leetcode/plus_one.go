package leetcode

// LC 66
func plusOne(digits []int) []int {
	n := len(digits)

	for index := n - 1; index >= 0; index-- {
		if digits[index] == 9 {
			digits[index] = 0
		} else {
			digits[index]++
			return digits
		}
	}

	digitsPlusOne := make([]int, n+1)
	digitsPlusOne[0] = 1

	return digitsPlusOne
}
