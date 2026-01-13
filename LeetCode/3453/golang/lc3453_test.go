package lc3453

import (
	"math"
	"testing"
)

func Test_separateSquares(t *testing.T) {
	testcases := []struct {
		squares  [][]int
		expected float64
	}{
		{[][]int{{0, 0, 1}, {2, 2, 1}}, 1.0},
		{[][]int{{0, 0, 2}, {1, 1, 1}}, 1.16667},
	}

	ACCEPTED_EPSILON := 0.00001
	for i, tc := range testcases {
		actual := separateSquares(tc.squares)
		if math.Abs(actual-tc.expected) > ACCEPTED_EPSILON {
			t.Errorf("Testcase %02d (%v) failed: want %.6f, got %.6f", i+1, tc.squares, tc.expected, actual)
		}
	}
}
