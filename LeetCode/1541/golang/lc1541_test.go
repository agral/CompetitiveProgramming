package lc1541

import "testing"

func Test_minInsertions(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"(()))", 1},   // insert one ')' at the end.
		{"())", 0},     // already balanced
		{"))())(", 3},  // insert '(' at the beginning; then "))" at the end.
		{"()", 1},      // insert ')' at the end.
		{")))))))", 5}, // WA #01, insert '(((()'.
	}

	for i, tc := range testcases {
		actual := minInsertions(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase minInsertions#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
