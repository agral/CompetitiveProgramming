package lc1390

import "testing"

func Test_sumFourDivisors(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{21, 4, 7}, 32 /* 21: 1, 3, 7, 21 */},
		{[]int{21, 21}, 64 /* 21: sum is 32, count twice */},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for i, tc := range testcases {
		actual := sumFourDivisors(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
