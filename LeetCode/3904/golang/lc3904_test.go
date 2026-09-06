package lc3904

import "testing"

func Test_firstStableIndex(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{5, 1, 0, 4}, 3, 3},
		{[]int{3, 2, 1}, 1, -1},
		{[]int{0}, 0, 0},
	}

	for i, tc := range testcases {
		actual := firstStableIndex(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase firstStableIndex#%02d (%v, %d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
