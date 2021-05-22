package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	expected := [][]string{{"bat"}, {"eat", "tea", "ate"}, {"tan", "nat"}}

	assert.ElementsMatch(t, got, expected, "Did not get expected result. Wanted %v, got %v\n",
		expected, got)
}
