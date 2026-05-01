package lc0396

import "testing"

func Test_maxRotateFunction(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{4, 3, 2, 6}, 26},
		{[]int{100}, 0},
	}

	for i, tc := range testcases {
		actual := maxRotateFunction(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase maxRotateFunction#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
