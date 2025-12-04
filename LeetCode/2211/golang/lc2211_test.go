package lc2211

import "testing"

func Test_countCollisions(t *testing.T) {
	testcases := []struct {
		directions string
		expected   int
	}{
		{"RLRSLL", 5},
		{"LLRR", 0},
	}

	for i, tc := range testcases {
		actual := countCollisions(tc.directions)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.directions, tc.expected, actual)
		}
	}
}
