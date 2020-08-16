package leetcode

// LC 11
func maxArea(height []int) int {
	maxArea, left, right := 0, 0, len(height)-1

	for left < right {
		width := right - left
		high := 0

		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}

		temp := width * high

		if temp > maxArea {
			maxArea = temp
		}
	}

	return maxArea
}
