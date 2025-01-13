package lc3223

import "testing"

func TestMinimumLength(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"abaacbcbb", 5},
		{"aa", 2},
	}

	for i, tc := range testcases {
		actual := minimumLength(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %d, got %d", i+1, tc.s, tc.expected, actual)
		}
	}
}
