package lc3362

import "testing"

func Test_maxRemoval(t *testing.T) {
	testcases := []struct {
		nums     []int
		queries  [][]int
		expected int
	}{
		{[]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}}, 1},
		{[]int{1, 1, 1, 1}, [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}, 2},
		{[]int{1, 2, 3, 4}, [][]int{{0, 3}}, -1},
	}

	for i, tc := range testcases {
		actual := maxRemoval(tc.nums, tc.queries)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.nums, tc.queries, tc.expected, actual)
		}
	}
}
