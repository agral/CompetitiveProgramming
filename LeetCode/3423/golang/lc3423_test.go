package lc3423

import "testing"

func Test_maxAdjacentDistance(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 4}, 3},
		{[]int{-5, -10, -5}, 5},
	}

	for i, tc := range testcases {
		actual := maxAdjacentDistance(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
