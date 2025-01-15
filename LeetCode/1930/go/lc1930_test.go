package lc1930

import "testing"

func TestCountPalindromicSubsequence(t *testing.T) {
	testingData := []struct {
		s        string
		expected int
	}{
		{"aabca", 3},   // aaa, aba, aca
		{"adc", 0},     // (none)
		{"bbcbaba", 4}, // "bbb", "bcb", "bab", "aba"
	}

	for i, d := range testingData {
		actual := countPalindromicSubsequence(d.s)
		if actual != d.expected {
			t.Errorf("Testcase %02d (%s) failed: want %d, got %d", i, d.s, d.expected, actual)
		}
	}
}
