package lc2942

import (
	"slices"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	testcases := []struct {
		words    []string
		x        byte
		expected []int
	}{
		{[]string{"leet", "code"}, 'e', []int{0, 1}},
		{[]string{"abc", "bcd", "aaaa", "cbc"}, 'a', []int{0, 2}},
		{[]string{"abc", "bcd", "aaaa", "cbc"}, 'z', []int{}},
	}

	for i, tc := range testcases {
		actual := findWordsContaining(tc.words, tc.x)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v | %c) failed: want %d, got %d",
				i+1, tc.words, tc.x, tc.expected, actual)
		}
	}
}
