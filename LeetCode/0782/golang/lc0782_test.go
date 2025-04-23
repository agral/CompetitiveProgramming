package lc0782

import "testing"

func Test_movesToChessboard(t *testing.T) {
	testcases := []struct {
		board    [][]int
		expected int
	}{
		{[][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}}, 2},
		{[][]int{{0, 1}, {1, 0}}, 0},
		{[][]int{{0, 1}, {0, 1}}, -1},
		{[][]int{{1, 0}, {1, 0}}, -1},
	}

	for i, tc := range testcases {
		actual := movesToChessboard(tc.board)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.board, tc.expected, actual)
		}
	}
}
