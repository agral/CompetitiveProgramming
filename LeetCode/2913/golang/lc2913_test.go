package lc2913

import "testing"

func Test_sumCounts(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 1}, 15},
		{[]int{1, 1}, 3},
	}

	for i, tc := range testcases {
		actual := sumCounts(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase sumCounts#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
