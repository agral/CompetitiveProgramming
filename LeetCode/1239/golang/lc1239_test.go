package lc1239

import "testing"

func Test_maxLength(t *testing.T) {
	testcases := []struct {
		arr      []string
		expected int
	}{
		{[]string{"un", "iq", "ue"}, 4},
		{[]string{"cha", "r", "act", "ers"}, 6},
		{[]string{"abcdefghijklmnopqrstuvwxyz"}, 26},
	}

	for i, tc := range testcases {
		actual := maxLength(tc.arr)
		if actual != tc.expected {
			t.Errorf("Testcase maxLength#%02d (%v) failed: want %d, got %d",
				i+1, tc.arr, tc.expected, actual)
		}
	}
}
