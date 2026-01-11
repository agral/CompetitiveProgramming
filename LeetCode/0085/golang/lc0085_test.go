package lc0085

import "testing"

func Test_maximalRectangle(t *testing.T) {
	testcases := []struct {
		matrix   [][]byte
		expected int
	}{
		{[][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}, 6},
		{[][]byte{
			{'0'},
		}, 0},
		{[][]byte{
			{'1'},
		}, 1},
	}

	for i, tc := range testcases {
		actual := maximalRectangle(tc.matrix)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.matrix, tc.expected, actual)
		}
	}
}
