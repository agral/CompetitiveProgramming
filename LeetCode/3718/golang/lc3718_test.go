package lc3718

import "testing"

func Test_missingMultiple(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{8, 2, 3, 4, 6}, 2, 10},
		{[]int{1, 4, 7, 10, 15}, 5, 5},
	}

	for i, tc := range testcases {
		actual := missingMultiple(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase missingMultiple#%02d (%v, k=%d) failed: want %d, got %d",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
