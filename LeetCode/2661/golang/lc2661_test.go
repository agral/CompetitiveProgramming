package lc2661

import "testing"

func TestFirstCompleteIndex(t *testing.T) {
	testcases := []struct {
		arr      []int
		mat      [][]int
		expected int
	}{
		{[]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}}, 2},
		{[]int{2, 8, 7, 4, 1, 3, 5, 6, 9}, [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}, 3},
	}

	for i, tc := range testcases {
		actual := firstCompleteIndex(tc.arr, tc.mat)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
