package lc0377

import "testing"

func TestCombinationSum4(t *testing.T) {
	testcases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
		{[]int{3}, 3, 1},
	}

	for i, tc := range testcases {
		actual := combinationSum4(tc.nums, tc.target)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %d) failed: want %d, got %d",
				i+1, tc.nums, tc.target, tc.expected, actual)
		}
	}
}
