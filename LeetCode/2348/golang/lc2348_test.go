package lc2348

import "testing"

func Test_zeroFilledSubarray(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int64
	}{
		{[]int{1, 3, 0, 0, 2, 0, 0, 4}, int64(6)},
		{[]int{0, 0, 0, 2, 0, 0}, int64(9)},
		{[]int{1, 2, 4, 5}, int64(0)},
	}

	for i, tc := range testcases {
		actual := zeroFilledSubarray(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
