package lc1590

import "testing"

func Test_minSubarray(t *testing.T) {
	testcases := []struct {
		nums     []int
		p        int
		expected int
	}{
		{[]int{3, 1, 4, 2}, 6, 1},
		{[]int{6, 3, 5, 2}, 9, 2},
		{[]int{1, 2, 3}, 3, 0},
		{[]int{1, 2, 3}, 7, -1},
	}

	for i, tc := range testcases {
		actual := minSubarray(tc.nums, tc.p)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.p, tc.expected, actual)
		}
	}
}
