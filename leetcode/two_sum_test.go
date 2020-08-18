package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}

	assert.ElementsMatch(t, got, expected, "Did not get expected result. Wanted %v, got %v\n",
		expected, got)
}
