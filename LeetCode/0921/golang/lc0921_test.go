package lc0921

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"())", 1},
		{"(((", 3},
		{"()()()", 0},
		{")", 1}, // yes, just insert a '(' in the beginning - we're allowed to.
	}

	for i, tc := range testcases {
		actual := minAddToMakeValid(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase minAddToMakeValid#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
