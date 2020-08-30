package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	got := productExceptSelf([]int{1, 2, 3, 4})
	expected := []int{24, 12, 8, 6}

	assert.ElementsMatch(t, got, expected, "Did not get expected result. Wanted %v, got %v\n",
		expected, got)
}
