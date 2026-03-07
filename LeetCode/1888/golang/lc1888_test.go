package lc1888

import "testing"

func Test_minFlips(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"111000", 2}, // 100011 -> 101010
		{"010", 0},    // already alternating
		{"1110", 1},   // 1110 -> 1010
	}

	for i, tc := range testcases {
		actual := minFlips(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase minFlips#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
