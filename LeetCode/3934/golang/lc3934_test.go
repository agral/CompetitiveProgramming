package lc3934

import "testing"

func Test_smallestUniqueSubarray(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 3, 3}, 3},
		{[]int{2, 1, 2, 3, 3}, 1},
		{[]int{1, 1, 2, 2, 1}, 2},
		{[]int{10531, 98357, 10531, 98357, 10531, 98357, 10531, 98357}, 6},
	}

	for i, tc := range testcases {
		actual := smallestUniqueSubarray(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase smallestUniqueSubarray#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
