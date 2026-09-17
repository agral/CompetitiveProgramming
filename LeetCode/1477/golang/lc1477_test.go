package lc1477

import "testing"

func Test_minSumOfLengths(t *testing.T) {
	testcases := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{3, 2, 2, 4, 3}, 3, 2},
		{[]int{7, 3, 4, 7}, 7, 2},
		{[]int{4, 3, 2, 6, 2, 3, 4}, 6, -1},
	}

	for i, tc := range testcases {
		actual := minSumOfLengths(tc.arr, tc.target)
		if actual != tc.expected {
			t.Errorf("Testcase minSumOfLengths#%02d (arr=%v, target=%d) failed: want %d, got %d",
				i+1, tc.arr, tc.target, tc.expected, actual)
		}
	}
}
