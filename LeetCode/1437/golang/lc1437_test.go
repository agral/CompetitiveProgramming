package lc1437

import "testing"

func Test_kLengthApart(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{1, 0, 0, 0, 1, 0, 0, 1}, 2, true},
		{[]int{1, 0, 0, 1, 0, 1}, 2, false},
		{[]int{1, 0, 1, 0, 1, 0, 0, 1}, 1, true},
		{[]int{1, 0, 1, 0, 1, 0, 0, 1}, 0, true},
		{[]int{1, 0, 1, 0, 1, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1}, 2, true},
		{[]int{1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1}, 3, false},
	}

	for i, tc := range testcases {
		actual := kLengthApart(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | k=%d) failed: want %t, got %t",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
