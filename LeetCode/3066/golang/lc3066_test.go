package lc3066

import "testing"

func Test_minOperations(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{2, 11, 10, 1, 3}, 10, 2},
		{[]int{1, 1, 2, 4, 9}, 20, 4},
	}

	for i, tc := range testcases {
		actual := minOperations(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d", i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
