package lc1291

import (
	"slices"
	"testing"
)

func Test_sequentialDigits(t *testing.T) {
	testcases := []struct {
		low      int
		high     int
		expected []int
	}{
		{100, 300, []int{123, 234}},
		{1000, 13000, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
	}

	for i, tc := range testcases {
		actual := sequentialDigits(tc.low, tc.high)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase sequentialDigits#%02d (low=%d, high=%d) failed: want %v, got %v",
				i+1, tc.low, tc.high, tc.expected, actual)
		}
	}
}
