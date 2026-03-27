package lc0932

import (
	"slices"
	"testing"
)

func Test_beautifulArray(t *testing.T) {
	testcases := []struct {
		n        int
		expected []int
	}{
		//{4, []int{2, 1, 4, 3}},
		{4, []int{3, 1, 2, 4}}, // also a valid one

		//{5, []int{3, 1, 2, 5, 4}},
		{5, []int{3, 5, 1, 2, 4}},
	}

	for i, tc := range testcases {
		actual := beautifulArray(tc.n)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase beautifulArray#%02d (%d) failed: want %v, got %v",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
