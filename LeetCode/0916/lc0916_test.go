package lc0916

import (
	"slices"
	"testing"
)

func TestWordSubsets(t *testing.T) {
	testcases := []struct {
		words1   []string
		words2   []string
		expected []string
	}{
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"e", "o"},
			[]string{"facebook", "google", "leetcode"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"l", "e"},
			[]string{"apple", "google", "leetcode"},
		},
	}

	for i, tc := range testcases {
		actual := wordSubsets(tc.words1, tc.words2)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d failed: want %v, got %v", i+1, tc.expected, actual)
		}
	}
}
