package lc3360

import "testing"

func Test_canAliceWin(t *testing.T) {
	testcases := []struct {
		n        int
		expected bool
	}{
		{12, true},
		{1, false},
		{9, false},
		{10, true}, // 0 + 10, change
		{18, true},
		{19, false}, // 10 + 9, change
		{26, false},
		{27, true}, // 19 + 8, change
		{33, true},
		{34, false}, // 27 + 7, change
		// and so on. Let's leave this O(10) == O(1) algo, not turning this into one if statement.
	}

	for i, tc := range testcases {
		actual := canAliceWin(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase canAliceWin#%02d (%d) failed: want %t, got %t",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
