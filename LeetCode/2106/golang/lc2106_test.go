package lc2106

import "testing"

func Test_maxTotalFruits(t *testing.T) {
	testcases := []struct {
		fruits   [][]int
		startPos int
		k        int
		expected int
	}{
		{[][]int{{2, 8}, {6, 3}, {8, 6}}, 5, 4, 9},
		{[][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, 5, 4, 14},
		{[][]int{{0, 3}, {6, 4}, {8, 5}}, 3, 2, 0},
	}

	for i, tc := range testcases {
		actual := maxTotalFruits(tc.fruits, tc.startPos, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d | %d) failed: want %d, got %d",
				i+1, tc.fruits, tc.startPos, tc.k, tc.expected, actual)
		}
	}
}
