package lc1408

import (
	"slices"
	"testing"
)

func TestStringMatching(t *testing.T) {
	testcases := []struct {
		words    []string
		expected []string
	}{
		{[]string{"mass", "as", "hero", "superhero"}, []string{"as", "hero"}},
		{[]string{"leetcode", "et", "code"}, []string{"et", "code"}},
		{[]string{"blue", "green", "bu"}, []string{}},
		{[]string{"leetcoder", "leetcode", "od", "hamlet", "am"}, []string{"leetcode", "od", "am"}},
	}
	for i, tc := range testcases {
		actual := stringMatching(tc.words)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v", i, tc.words, tc.expected, actual)
		}
	}
}
