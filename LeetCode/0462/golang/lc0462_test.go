package lc0462

import "testing"

func Test_minMoves2(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 10, 2, 9}, 16},
	}

	for i, tc := range testcases {
		actual := minMoves2(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase minMoves2#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
