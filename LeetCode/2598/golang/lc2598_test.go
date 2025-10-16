package lc2598

import "testing"

func Test_findSmallestInteger(t *testing.T) {
	testcases := []struct {
		nums     []int
		value    int
		expected int
	}{
		{[]int{1, -10, 7, 13, 6, 8}, 5, 4},
		{[]int{1, -10, 7, 13, 6, 8}, 7, 2},
	}

	for i, tc := range testcases {
		actual := findSmallestInteger(tc.nums, tc.value)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.value, tc.expected, actual)
		}
	}
}
