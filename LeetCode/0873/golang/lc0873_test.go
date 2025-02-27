package lc0873

import "testing"

func Test_lenLongestFibSubseq(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 5},  // [1, 2, 3, 5, 8]
		{[]int{1, 3, 7, 11, 12, 14, 18}, 3}, // [1, 11, 12], [3, 11, 14], [7, 11, 18].
	}

	for i, tc := range testcases {
		actual := lenLongestFibSubseq(tc.arr)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.arr, tc.expected, actual)
		}
	}
}
