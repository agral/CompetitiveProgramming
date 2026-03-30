package lc2840

import "testing"

func Test_checkStrings(t *testing.T) {
	testcases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"abcdba", "cabdab", true},
		{"abe", "bea", false},
	}

	for i, tc := range testcases {
		actual := checkStrings(tc.s1, tc.s2)
		if actual != tc.expected {
			t.Errorf("Testcase checkStrings#%02d (%s | %s) failed: want %t, got %t",
				i+1, tc.s1, tc.s2, tc.expected, actual)
		}
	}
}
