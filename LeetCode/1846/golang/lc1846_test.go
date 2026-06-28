package lc1846

import "testing"

func Test_maximumElementAfterDecrementingAndRearranging(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected int
	}{
		{[]int{2, 2, 1, 2, 1}, 2},
		{[]int{100, 1, 1000}, 3},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for i, tc := range testcases {
		actual := maximumElementAfterDecrementingAndRearranging(tc.arr)
		if actual != tc.expected {
			t.Errorf("Testcase maximumElementAfterDecrementingAndRearranging#%02d (%v) failed: want %d, got %d",
				i+1, tc.arr, tc.expected, actual)
		}
	}
}
