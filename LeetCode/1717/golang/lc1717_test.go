package lc1717

import "testing"

func Test_maximumGain(t *testing.T) {
	testcases := []struct {
		s        string
		x        int
		y        int
		expected int
	}{
		{"cdbcbbaaabab", 4, 5, 19},
		{"aabbaaxybbaabb", 5, 4, 20},
	}

	for i, tc := range testcases {
		actual := maximumGain(tc.s, tc.x, tc.y)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s | ab=%d | ba=%d) failed: want %d, got %d",
				i+1, tc.s, tc.x, tc.y, tc.expected, actual)
		}
	}
}
