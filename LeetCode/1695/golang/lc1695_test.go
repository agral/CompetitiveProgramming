package lc1695

import "testing"

func Test_maximumUniqueSubarray(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{4, 2, 4, 5, 6}, 17 /*{2, 4, 5, 6}*/},
		{[]int{5, 2, 1, 2, 5, 2, 1, 2, 5}, 8 /*{5, 2, 1}*/},
	}

	for i, tc := range testcases {
		actual := maximumUniqueSubarray(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
