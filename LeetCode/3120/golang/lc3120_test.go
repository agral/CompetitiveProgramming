package lc3120

import "testing"

func Test_numberOfSpecialChars(t *testing.T) {
	testcases := []struct {
		word     string
		expected int
	}{
		{"aaAbcBC", 3},
		{"abc", 0},
		{"abBCab", 1},
	}

	for i, tc := range testcases {
		actual := numberOfSpecialChars(tc.word)
		if actual != tc.expected {
			t.Errorf("Testcase numberOfSpecialChars#%02d (%v) failed: want %d, got %d",
				i+1, tc.word, tc.expected, actual)
		}
	}
}
