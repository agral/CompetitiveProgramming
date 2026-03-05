package lc1758

import "testing"

func Test_minOperations(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"0100", 1},
		{"10", 0},
		{"1111", 2},
	}

	for i, tc := range testcases {
		actual := minOperations(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase minOperations#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
