package lc1128

import "testing"

func Test_numEquivDominoPairs(t *testing.T) {
	testcases := []struct {
		dominoes [][]int
		expected int
	}{
		{[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 1}, {2, 1}, {2, 2}}, 3},
		{[][]int{{1, 1}, {1, 1}, {2, 3}, {5, 6}}, 1},          // one pair of [1, 1]
		{[][]int{{1, 3}, {2, 1}, {3, 1}, {5, 6}, {1, 3}}, 3},  // three instances of [1, 3] => 3 pairs
		{[][]int{{1, 3}, {3, 1}, {3, 1}, {1, 3}}, 6},          // four instances of [1, 3] => 6 pairs
		{[][]int{{1, 3}, {3, 1}, {3, 1}, {1, 3}, {3, 1}}, 10}, // five instances of [1, 3] => 10 pairs
		{[][]int{{1, 3}, {2, 5}, {3, 1}, {1, 3}, {5, 2}}, 4},  // 3 * [1, 3], 2 * [2, 5] => 4 pairs
	}

	for i, tc := range testcases {
		actual := numEquivDominoPairs(tc.dominoes)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.dominoes, tc.expected, actual)
		}
	}
}
