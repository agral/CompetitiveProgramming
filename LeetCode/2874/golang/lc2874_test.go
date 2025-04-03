package lc2874

import "testing"

func Test_maximumTripletValue(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int64
	}{
		{[]int{12, 6, 1, 2, 7}, 77},
		{[]int{1, 10, 3, 4, 19}, 133},
		{[]int{1, 2, 3}, 0},
	}

	for i, tc := range testcases {
		actual := maximumTripletValue(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.nums, tc.expected, actual)
		}
	}
}
