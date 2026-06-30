package lc1358

import "testing"

func Test_numberOfSubstrings(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"abcabc", 10},
		{"aaacb", 3},
		{"abc", 1},
	}

	for i, tc := range testcases {
		actual := numberOfSubstrings(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase numberOfSubstrings#%02d (%s) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
