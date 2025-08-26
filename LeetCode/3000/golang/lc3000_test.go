package lc3000

import "testing"

func Test_areaOfMaxDiagonal(t *testing.T) {
	testcases := []struct {
		dimensions [][]int
		expected   int
	}{
		{[][]int{{9, 3}, {8, 6}}, 48},
		{[][]int{{3, 4}, {4, 3}}, 12},
	}

	for i, tc := range testcases {
		actual := areaOfMaxDiagonal(tc.dimensions)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.dimensions, tc.expected, actual)
		}
	}
}
