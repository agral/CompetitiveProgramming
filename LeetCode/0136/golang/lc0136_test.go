package lc0136

import "testing"

func Test_singleNumber(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for i, tc := range testcases {
		actual := singleNumber(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase singleNumber#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
