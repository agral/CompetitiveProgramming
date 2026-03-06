package lc1784

import "testing"

func Test_checkOnesSegment(t *testing.T) {
	testcases := []struct {
		s        string
		expected bool
	}{
		{"1001", false},
		{"10101", false},
		{"110", true},
		{"1", true},
		{"10000000", true},
		{"11111110", true},
		{"11111111", true},
	}

	for i, tc := range testcases {
		actual := checkOnesSegment(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase checkOnesSegment#%02d (%s) failed: want %t, got %t",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
