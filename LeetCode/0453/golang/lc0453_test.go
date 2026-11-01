package lc0453

import "testing"

func Test_minMoves(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 1}, 0},
	}

	for i, tc := range testcases {
		actual := minMoves(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase minMoves#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
