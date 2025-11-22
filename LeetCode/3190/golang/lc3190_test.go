package lc3190

import "testing"

func Test_minimumOperations(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 3},
		{[]int{3, 6, 9, 12}, 0},
	}

	for i, tc := range testcases {
		actual := minimumOperations(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
