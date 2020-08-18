package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	got := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	bubbleSort(got)

	assert.ElementsMatch(t, got, expected, "Did not get expected result. Wanted %v, got %v\n",
		expected, got)
}
