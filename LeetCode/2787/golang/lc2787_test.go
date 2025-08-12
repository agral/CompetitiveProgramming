package lc2787

import "testing"

func Test_numberOfWays(t *testing.T) {
	testcases := []struct {
		n        int
		x        int
		expected int
	}{
		{10, 2, 1},
		{4, 1, 2},
	}

	for i, tc := range testcases {
		actual := numberOfWays(tc.n, tc.x)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
