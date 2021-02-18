package leetcode

// LC 231
func powerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}
