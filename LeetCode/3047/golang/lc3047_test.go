package lc3047

import "testing"

func Test_largestSquareArea(t *testing.T) {
	testcases := []struct {
		bottomLeft [][]int
		topRight   [][]int
		expected   int64
	}{
		{[][]int{{1, 1}, {2, 2}, {3, 1}}, [][]int{{3, 3}, {4, 4}, {6, 6}}, int64(1)},
		{[][]int{{1, 1}, {1, 3}, {1, 5}}, [][]int{{5, 5}, {5, 7}, {5, 9}}, int64(4)},
		{[][]int{{1, 1}, {2, 2}, {1, 2}}, [][]int{{3, 3}, {4, 4}, {3, 4}}, int64(1)},
		{[][]int{{1, 1}, {3, 3}, {3, 1}}, [][]int{{2, 2}, {4, 4}, {4, 2}}, int64(0)},
	}

	for i, tc := range testcases {
		actual := largestSquareArea(tc.bottomLeft, tc.topRight)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.bottomLeft, tc.topRight, tc.expected, actual)
		}
	}
}
