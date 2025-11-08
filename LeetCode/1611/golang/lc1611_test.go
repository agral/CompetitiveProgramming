package lc1611

import "testing"

func Test_minimumOneBitOperations(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{3, 2},
		{6, 4},
	}

	for i, tc := range testcases {
		actual := minimumOneBitOperations(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
