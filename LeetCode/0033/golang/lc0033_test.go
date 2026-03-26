package lc0033

import "testing"

func Test_search(t *testing.T) {
	testcases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}

	for i, tc := range testcases {
		actual := search(tc.nums, tc.target)
		if actual != tc.expected {
			t.Errorf("Testcase search#%02d (%v | target=%d) failed: want %d, got %d",
				i+1, tc.nums, tc.target, tc.expected, actual)
		}
	}
}
