package lc3660

import (
	"slices"
	"testing"
)

func Test_maxValue(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{2, 1, 3}, []int{2, 2, 3}},
		{[]int{2, 3, 1}, []int{3, 3, 3}},
	}

	for i, tc := range testcases {
		actual := maxValue(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase maxValue#%02d (%v) failed. want:\n%v,\n\ngot:\n%v\n\n",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
