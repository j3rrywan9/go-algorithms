package interview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookAndSay(t *testing.T) {
	expected := "111221"

	got := lookAndSay(1, 4)

	assert.Equal(t, got, expected, "Did not get expected result. Wanted %v, got %v\n",
		expected, got)
}
