package lc2785

import "testing"

func Test_sortVowels(t *testing.T) {
	testcases := []struct {
		s        string
		expected string
	}{
		{"lEetcOde", "lEOtcede"},
		{"lYmpH", "lYmpH"},
	}

	for i, tc := range testcases {
		actual := sortVowels(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %s, got %s", i+1, tc.s, tc.expected, actual)
		}
	}
}
