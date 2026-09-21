package lc3867

import "testing"

func Test_gcdSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int64
	}{
		{[]int{2, 6, 4}, 2},
		{[]int{3, 6, 2, 8}, 5},
	}

	for i, tc := range testcases {
		actual := gcdSum(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase gcdSum#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
