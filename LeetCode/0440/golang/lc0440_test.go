package lc0440

import (
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	testcases := []struct {
		n        int
		k        int
		expected int
	}{
		{13, 2, 10}, // 1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9
		{13, 3, 11}, // 1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9
		{13, 6, 2},  // 1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9
		{2, 2, 2},
	}

	for i, tc := range testcases {
		actual := findKthNumber(tc.n, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d, %d) failed: want %d, got %d", i+1, tc.n, tc.k, tc.expected, actual)
		}
	}
}
