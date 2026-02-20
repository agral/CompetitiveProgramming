package lc0761

import "testing"

func Test_makeLargestSpecial(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{"11011000", "11100100"},
		{"10", "10"},
	}

	for i, tc := range testcases {
		actual := makeLargestSpecial(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase makeLargestSpecial#%02d (%v) failed: want %s, got %s",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
