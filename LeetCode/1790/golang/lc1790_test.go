package lc1790

import "testing"

func Test_areAlmostEqual(t *testing.T) {
	testcases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"bank", "kanb", true},
		{"attack", "defend", false},
		{"kelb", "kelb", true},
		{"aabaa", "aacaa", false},
		{"aabca", "aacda", false},
	}

	for i, tc := range testcases {
		actual := areAlmostEqual(tc.s1, tc.s2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s | %s) failed: want %t, got %t",
				i+1, tc.s1, tc.s2, tc.expected, actual)
		}
	}
}
