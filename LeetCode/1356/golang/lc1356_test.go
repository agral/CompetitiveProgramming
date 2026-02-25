package lc1356

import (
	"slices"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			[]int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		{
			[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
	}

	for i, tc := range testcases {
		actual := sortByBits(tc.arr)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase sortByBits#%02d (%v) failed: want %v, got %v",
				i+1, tc.arr, tc.expected, actual)
		}
	}
}
