package lc3069

import (
	"slices"
	"testing"
)

func Test_resultArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{5, 4, 3, 8}, []int{5, 3, 4, 8}},
	}

	for i, tc := range testcases {
		actual := resultArray(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase resultArray#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
