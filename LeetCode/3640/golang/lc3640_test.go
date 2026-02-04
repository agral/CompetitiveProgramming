package lc3640

import "testing"

func Test_maxSumTrionic(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int64
	}{
		{[]int{0, -2, -1, -3, 0, 2, -1}, -4}, // example #01
		{[]int{1, 4, 2, 7}, 14},              // example #02

		// A long trionic slice. Check that it adds all entries.
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3, 4, 5}, 39},

		// short trionic, then long trionic. Check that it records the correct long trionic seq.
		{[]int{1, 5, 1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3, 4, 5}, 39},

		// long trionic, then short trionic. Check that it records the correct long trionic seq.
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3, 4, 5, 1, 5}, 39},

		// short trionic with bigger sum, then long trionic with lesser sum.
		// Check that it records the correct short trionic seq with bigger sum.
		{[]int{1, 101, -11, 2, 1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3, 4, 5}, 93},

		// a nice big sum, but it's not trionic - numbers are repeated.
		// a small sum at the end must be returned - the only trionic entry.
		{[]int{100, 101, 102, 101, 100, 100, 1, 3, 2, 4}, 10},

		// a nice big sum, but it's not trionic - numbers are repeated.
		// a small sum from the trionic sequence in the beginning must be returned.
		{[]int{2, 4, 3, 6, 6, 101, 102, 101, 100, 100, 1, 3, 2, 4}, 15},

		// Trionic all the way, but the valid answer should use only positives -
		// - that is, can the solution work out which numbers *not* to use?:
		// here only a subarray {1, 2, 3, 2, 1, 2, 3} should be used.
		{[]int{-3, -2, -1, 1, 2, 3, 2, 1, 2, 3}, 14},

		{[]int{1, 2, 3, 2, 1, 2, 3, -1, -2, -3, 100}, 100},

		{[]int{-273, 85, -636, -486, -374, 331}, -1310},
		{[]int{2, 993, -791, -635, -569}, -431},
		{[]int{-754, 167, -172, 202, 735, -941, 992}, 988},
		{[]int{-222, 209, -246, 680, -662, 412}, 421},
	}

	for i, tc := range testcases {
		actual := maxSumTrionic(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase maxSumTrionic#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
