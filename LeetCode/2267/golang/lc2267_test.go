package lc2267

import "testing"

func Test_hasValidPath(t *testing.T) {
	testcases := []struct {
		grid     [][]byte
		expected bool
	}{
		{[][]byte{
			{'(', '(', '('},
			{')', '(', ')'},
			{'(', '(', ')'},
			{'(', '(', ')'},
		}, true},

		{[][]byte{
			{')', ')'},
			{'(', '('},
		}, false},
	}

	for i, tc := range testcases {
		actual := hasValidPath(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase hasValidPath#%02d (%v) failed: want %t, got %t",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
