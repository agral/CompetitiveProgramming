package lc2900

import (
	"slices"
	"testing"
)

func Test_getLongestSubsequence(t *testing.T) {
	testcases := []struct {
		words    []string
		groups   []int
		expected []string
	}{
		{[]string{"e", "a", "b"}, []int{0, 0, 1}, []string{"e", "b"}},
		{[]string{"a", "b", "c", "d"}, []int{1, 0, 1, 1}, []string{"a", "b", "c"}},
	}

	for i, tc := range testcases {
		actual := getLongestSubsequence(tc.words, tc.groups)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v | %v) failed: want %v, got %v",
				i+1, tc.words, tc.groups, tc.expected, actual)
		}
	}
}
