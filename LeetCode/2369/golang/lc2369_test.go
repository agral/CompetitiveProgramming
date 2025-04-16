package lc2369

import "testing"

func Test_validPartition(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{4, 4, 4, 5, 6}, true}, // valid partitioning: [4, 4], [4, 5, 6]
		{[]int{1, 1, 1, 2}, false},
		{[]int{1, 2, 3, 4, 5, 6}, true}, // valid p: [1, 2, 3], [4, 5, 6]
		{[]int{1, 2, 3, 4, 5, 6, 7}, false},
		{[]int{1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4}, true},
	}

	for i, tc := range testcases {
		actual := validPartition(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.nums, tc.expected, actual)
		}
	}
}
