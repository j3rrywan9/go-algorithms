package leetcode

// LC 20
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)

	for _, r := range s {
		if r == '[' || r == '(' || r == '{' {
			stack = append(stack, r)
		} else if (r == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			(r == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			(r == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
