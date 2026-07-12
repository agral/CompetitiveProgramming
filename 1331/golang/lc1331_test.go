package lc1331

import (
	"slices"
	"testing"
)

func Test_arrayRankTransform(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected []int
	}{
		{[]int{40, 10, 20, 30}, []int{4, 1, 2, 3}},
		{[]int{101, 101, 101}, []int{1, 1, 1}},
		{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}, []int{5, 3, 4, 2, 8, 6, 7, 1, 3}},
		{[]int{}, []int{}}, // RE #01
	}

	for i, tc := range testcases {
		actual := arrayRankTransform(tc.arr)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase arrayRankTransform#%02d (%v) failed: want %v, got %v",
				i+1, tc.arr, tc.expected, actual)
		}
	}
}
