package lc2965

import (
	"slices"
	"testing"
)

func Test_findMissingAndRepeatedValues(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected []int
	}{
		{[][]int{{1, 3}, {2, 2}}, []int{2, 4}},
		{[][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}, []int{9, 5}},
	}

	for i, tc := range testcases {
		actual := findMissingAndRepeatedValues(tc.grid)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v", i+1, tc.grid, tc.expected, actual)
		}
	}
}
