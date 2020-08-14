package leetcode

// LC 283
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	anchor := 0

	for explorer := 0; explorer < len(nums); explorer++ {
		if nums[explorer] != 0 {
			if explorer != anchor {
				nums[explorer], nums[anchor] = nums[anchor], nums[explorer]
			}
			anchor++
		}
	}
}
