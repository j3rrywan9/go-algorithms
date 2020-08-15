package leetcode

// LC 1
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if _, found := m[complement]; found {
			return []int{m[complement], i}
		}

		m[nums[i]] = i
	}

	return nil
}
