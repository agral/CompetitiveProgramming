package lc1524

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 3, 5}, 4},
		{[]int{2, 4, 6}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 16},
	}

	for i, tc := range testcases {
		actual := numOfSubarrays(tc.arr)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.arr, tc.expected, actual)
		}
	}
}
