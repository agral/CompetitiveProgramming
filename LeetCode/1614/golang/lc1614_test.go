package lc1614

import "testing"

func Test_maxDepth(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
		{"()(())((()()))", 3},
	}

	for i, tc := range testcases {
		actual := maxDepth(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase maxDepth#%02d (%s) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
