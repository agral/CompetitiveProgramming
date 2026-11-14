package lc3364

import "testing"

func Test_minimumSumSubarray(t *testing.T) {
	testcases := []struct {
		nums     []int
		l        int
		r        int
		expected int
	}{
		{[]int{3, -2, 1, 4}, 2, 3, 1},
		{[]int{-2, 2, -3, 1}, 2, 3, -1},
		{[]int{1, 2, 3, 4}, 2, 3, 3},
	}

	for i, tc := range testcases {
		actual := minimumSumSubarray(tc.nums, tc.l, tc.r)
		if actual != tc.expected {
			t.Errorf("Testcase minimumSumSubarray#%02d (%v, l=%d, r=%d) failed: want %d, got %d",
				i+1, tc.nums, tc.l, tc.r, tc.expected, actual)
		}
	}
}
