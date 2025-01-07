package lc0287

import "testing"

func TestFindDuplicate(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 13, 14, 15, 16}, 13},
	}

	for i, tc := range testcases {
		actual := findDuplicate(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) has failed: want %d, got %d", i, tc.nums, tc.expected, actual)
		}
	}
}
