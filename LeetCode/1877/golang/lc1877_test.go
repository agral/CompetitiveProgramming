package lc1877

import "testing"

func Test_minPairSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 5, 2, 3}, 7},
		{[]int{3, 5, 4, 2, 4, 6}, 8},
	}

	for i, tc := range testcases {
		actual := minPairSum(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
