package lc0712

import "testing"

func Test_minimumDeleteSum(t *testing.T) {
	testcases := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"sea", "eat", 231},
		{"delete", "leet", 403},
	}

	for i, tc := range testcases {
		actual := minimumDeleteSum(tc.s1, tc.s2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s | %s) failed: want %d, got %d",
				i+1, tc.s1, tc.s2, tc.expected, actual)
		}
	}
}
