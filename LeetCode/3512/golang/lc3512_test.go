package lc3512

import "testing"

func Test_minOperations(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{3, 9, 7}, 5, 4},
		{[]int{4, 1, 3}, 4, 0},
		{[]int{3, 2}, 6, 5},
	}

	for i, tc := range testcases {
		actual := minOperations(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
