package lc2185

import "testing"

func TestPrefixCount(t *testing.T) {
	testcases := []struct {
		words    []string
		pref     string
		expected int
	}{
		{[]string{"pay", "attention", "practice", "attend"}, "at", 2},
		{[]string{"leetcode", "win", "loops", "success"}, "code", 0},
	}

	for i, tc := range testcases {
		actual := prefixCount(tc.words, tc.pref)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v, %s) failed: want %d, got %d", i, tc.words, tc.pref, tc.expected, actual)
		}
	}
}

