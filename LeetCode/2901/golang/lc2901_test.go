package lc2901

import (
	"slices"
	"testing"
)

func Test_getWordsInLongestSubsequence(t *testing.T) {
	testcases := []struct {
		words    []string
		groups   []int
		expected []string
	}{
		{[]string{"bab", "dab", "cab"}, []int{1, 2, 2}, []string{"bab", "dab"}},
		{[]string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}, []string{"a", "b", "c", "d"}},
	}

	for i, tc := range testcases {
		actual := getWordsInLongestSubsequence(tc.words, tc.groups)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v | %v) failed: want %v, got %v",
				i+1, tc.words, tc.groups, tc.expected, actual)
		}
	}
}
