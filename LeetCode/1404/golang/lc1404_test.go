package lc1404

import "testing"

func Test_numSteps(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"1101", 6}, // 1101 -> 1110 -> 111 -> 1000 -> 100 -> 10 -> 1; 6 steps
		{"10", 1},   // 10 -> 1; 1 step
		{"1", 0},    // corner case; already a 1.
		{"1001", 7}, // 1001 -> 1010 -> 101 -> 110 -> 11 -> 100 -> 10 -> 1; 7 steps
	}

	for i, tc := range testcases {
		actual := numSteps(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase numSteps#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
