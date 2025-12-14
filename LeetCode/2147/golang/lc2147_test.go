package lc2147

import "testing"

func Test_numberOfWays(t *testing.T) {
	testcases := []struct {
		corridor string
		expected int
	}{
		{"SSPPSPS", 3},
		{"PPSPSP", 1},
		{"S", 0},
	}

	for i, tc := range testcases {
		actual := numberOfWays(tc.corridor)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %d, got %d", i+1, tc.corridor, tc.expected, actual)
		}
	}
}
