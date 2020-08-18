package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	got := binarySearch([]int{2, 3, 6, 7, 9, 11, 14}, 11)
	expected := 5

	assert.Equal(t, got, expected, "Did not get expected result. Wanted %v, got %v\n", expected, got)
}
