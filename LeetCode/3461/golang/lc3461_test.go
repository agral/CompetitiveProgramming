package lc3461

import "testing"

func Test_hasSameDigits(t *testing.T) {
	testcases := []struct {
		s        string
		expected bool
	}{
		{"3902", true},
		{"34789", false},
		{"12345", false},
	}

	for i, tc := range testcases {
		actual := hasSameDigits(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.s, tc.expected, actual)
		}
	}
}
