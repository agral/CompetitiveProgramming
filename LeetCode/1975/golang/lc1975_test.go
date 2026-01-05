package lc1975

import "testing"

func Test_maxMatrixSum(t *testing.T) {
	testcases := []struct {
		matrix   [][]int
		expected int64
	}{
		{[][]int{
			{1, -1},
			{-1, 1},
		}, 4},
		{[][]int{
			{1, 2, 3},
			{-1, -2, -3},
			{1, 2, 3},
		}, 16},
	}

	for i, tc := range testcases {
		actual := maxMatrixSum(tc.matrix)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.matrix, tc.expected, actual)
		}
	}
}
