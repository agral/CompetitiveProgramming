package lc2294

import "testing"

func Test_partitionArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 6, 1, 2, 5}, 2, 2},
		{[]int{1, 2, 3}, 1, 2},
		{[]int{2, 2, 4, 5}, 0, 3},
	}

	for i, tc := range testcases {
		actual := partitionArray(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
