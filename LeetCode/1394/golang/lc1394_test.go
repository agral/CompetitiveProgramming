package lc1394

import "testing"

func Test_findLucky(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 2, 3, 4}, 2},
		{[]int{1, 2, 2, 3, 3, 3}, 3},
		{[]int{2, 2, 2, 3, 3}, -1},
	}

	for i, tc := range testcases {
		actual := findLucky(tc.arr)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.arr, tc.expected, actual)
		}
	}
}
