package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	got := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}

	moveZeroes(got)

	assert.ElementsMatch(t, got, expected, "Did not get expected result. Wanted %v, got %v\n",
		expected, got)
}
