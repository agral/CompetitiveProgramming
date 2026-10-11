package lc2778

import "testing"

func Test_sumOfSquares(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 21},
		{[]int{2, 7, 1, 19, 18, 3}, 63},
	}

	for i, tc := range testcases {
		actual := sumOfSquares(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase sumOfSquares#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
