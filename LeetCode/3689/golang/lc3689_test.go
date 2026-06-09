package lc3689

import "testing"

func Test_maxTotalValue(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int64
	}{
		{[]int{1, 3, 2}, 2, 4},
		{[]int{4, 2, 5, 1}, 3, 12},
	}

	for i, tc := range testcases {
		actual := maxTotalValue(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase maxTotalValue#%02d (%v | k=%d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
