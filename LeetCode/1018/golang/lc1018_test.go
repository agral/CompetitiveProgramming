package lc1018

import (
	"slices"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected []bool
	}{
		{[]int{0, 1, 1}, []bool{true, false, false}},
		{[]int{1, 1, 1}, []bool{false, false, false}},
	}

	for i, tc := range testcases {
		actual := prefixesDivBy5(tc.nums)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed.\nwant: %v,\n got: %v",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
