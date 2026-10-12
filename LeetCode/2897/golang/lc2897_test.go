package lc2897

import "testing"

func Test_maxSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{2, 6, 5, 8}, 2, 261},
		{[]int{4, 5, 4, 7}, 3, 90},
	}

	for i, tc := range testcases {
		actual := maxSum(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase maxSum#%02d (%v, k=%d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
