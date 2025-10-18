package lc3397

import "testing"

func Test_maxDistinctElements(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 2, 2, 3, 3, 4}, 2, 6}, // e.g. {-1, 0, 1, 2, 3, 4}
		{[]int{4, 4, 4, 4}, 1, 3},       // e.g. {3, 4, 4, 5}
	}

	for i, tc := range testcases {
		actual := maxDistinctElements(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
