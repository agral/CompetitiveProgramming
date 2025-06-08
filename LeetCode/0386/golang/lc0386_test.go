package lc0386

import (
	"slices"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	testcases := []struct {
		n        int
		expected []int
	}{
		{13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		{2, []int{1, 2}},
	}

	for i, tc := range testcases {
		actual := lexicalOrder(tc.n)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%d) failed: want %v, got %v", i+1, tc.n, tc.expected, actual)
		}
	}
}
