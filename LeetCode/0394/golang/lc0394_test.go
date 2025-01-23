package lc0394

import "testing"

func TestDecodeString(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	for i, tc := range testcases {
		actual := decodeString(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %q, got %q", i+1, tc.input, tc.expected, actual)
		}
	}
}
