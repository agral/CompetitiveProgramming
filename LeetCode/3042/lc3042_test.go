package lc3042

import "testing"

func TestCountPrefixSuffixPairs(t *testing.T) {
	testcases := []struct {
		words    []string
		expected int
	}{
		{[]string{"a", "aba", "ababa", "aa"}, 4},
		{[]string{"pa", "papa", "ma", "mama"}, 2},
		{[]string{"abab", "ab"}, 0},
	}

	for i, tc := range testcases {
		actual := countPrefixSuffixPairs(tc.words)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %d, got %d", i, tc.words, tc.expected, actual)
		}
	}
}
