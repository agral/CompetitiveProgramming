package lc1967

import "testing"

func Test_numOfStrings(t *testing.T) {
	testcases := []struct {
		patterns []string
		word     string
		expected int
	}{
		{[]string{"a", "abc", "bc", "d"}, "abc", 3},
		{[]string{"a", "b", "c"}, "aaaaabbbbb", 2},
		{[]string{"a", "a", "a"}, "ab", 3},
	}

	for i, tc := range testcases {
		actual := numOfStrings(tc.patterns, tc.word)
		if actual != tc.expected {
			t.Errorf("Testcase numOfStrings#%02d (%v | word=%s) failed: want %d, got %d",
				i+1, tc.patterns, tc.word, tc.expected, actual)
		}
	}
}
