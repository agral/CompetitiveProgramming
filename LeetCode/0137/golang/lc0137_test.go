package lc0137

import "testing"

func Test_singleNumber(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
		{[]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}, -4},
	}

	for i, tc := range testcases {
		actual := singleNumber(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase singleNumber#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
