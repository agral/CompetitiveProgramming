package lc1726

import "testing"

func Test_tupleSameProduct(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 4, 6}, 8},      // 2*6 == 3*4, combined a total of 8 times.
		{[]int{1, 2, 4, 5, 10}, 16}, // 1 * 10 == 2 * 5, 2 * 10 == 4 * 5; both combined a total of 8 times.
	}

	for i, tc := range testcases {
		actual := tupleSameProduct(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
