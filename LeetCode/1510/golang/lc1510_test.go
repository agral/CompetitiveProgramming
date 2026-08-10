package lc1510

import "testing"

func Test_winnerSquareGame(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, true},
		{5, false},
		{6, true},
		{7, false},
		{8, true},   // it's NOT-7, so must be winnable.
		{9, true},   // 3^2, an easy win.
		{10, false}, // 3^2+1, an easy lose. Also removing 2^2==4 does not help (6 is true, so not-6 is losing).
	}

	for i, tc := range testcases {
		actual := winnerSquareGame(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase winnerSquareGame#%02d (%d) failed: want %t, got %t",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
