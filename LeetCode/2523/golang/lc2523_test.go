package lc2523

import (
	"slices"
	"testing"
)

func Test_closestPrimes(t *testing.T) {
	testcases := []struct {
		left     int
		right    int
		expected []int
	}{
		{10, 19, []int{11, 13}},
		{4, 6, []int{-1, -1}},
	}

	for i, tc := range testcases {
		actual := closestPrimes(tc.left, tc.right)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%d, %d) failed: want %v, got %v",
				i+1, tc.left, tc.right, tc.expected, actual)
		}
	}
}
