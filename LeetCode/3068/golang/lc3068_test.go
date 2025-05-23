package lc3068

import "testing"

func Test_maximumValueSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		edges    [][]int
		expected int64
	}{
		{[]int{1, 2, 1}, 3, [][]int{{}}, 6},
		{[]int{2, 3}, 7, [][]int{{}}, 9},
		{[]int{7, 7, 7, 7, 7, 7}, 3, [][]int{{}}, 42},
	}

	for i, tc := range testcases {
		actual := maximumValueSum(tc.nums, tc.k, tc.edges)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
