package lc2273

import (
	"slices"
	"testing"
)

func Test_removeAnagrams(t *testing.T) {
	testcases := []struct {
		words    []string
		expected []string
	}{
		{[]string{"abba", "baba", "bbaa", "cd", "cd"}, []string{"abba", "cd"}},
		{[]string{"a", "b", "c", "d", "e"}, []string{"a", "b", "c", "d", "e"}},
		{[]string{"abba", "baba", "bbaa"}, []string{"abba"}},
	}

	for i, tc := range testcases {
		actual := removeAnagrams(tc.words)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v",
				i+1, tc.words, tc.expected, actual)
		}
	}
}
