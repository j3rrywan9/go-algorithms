package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Did not get expected result. Wanted %v, got %v\n", expected, got)
	}
}
