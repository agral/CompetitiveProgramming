package lc0089

import (
	"slices"
	"testing"
)

func Test_grayCode(t *testing.T) {
	testcases := []struct {
		n        int
		expected []int
	}{
		{1, []int{0, 1}},
		{2, []int{0, 1, 3, 2}},
		{3, []int{0, 1, 3, 2, 6, 7, 5, 4}},
	}

	for i, tc := range testcases {
		actual := grayCode(tc.n)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase grayCode#%02d (%d) failed: want %v, got %v",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
