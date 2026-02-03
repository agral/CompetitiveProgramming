package lc3637

import "testing"

func Test_isTrionic(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 3, 5, 4, 2, 6}, true}, // example #01
		{[]int{2, 1, 3}, false},         // example #02

		{[]int{1, 4, 3, 5}, true},          // a minimal trionic slice. Note: has 4 elements.
		{[]int{1, 2, 3, 2, 1}, false},      // duo-tonic, goes up and down. Not trionic.
		{[]int{1, 2, 3, 4, 5}, false},      // uno-tonic, goes only up. Not trionic.
		{[]int{5, 4, 3, 2, 1}, false},      // uno-tonic, goes only down. Not trionic.
		{[]int{1, 1, 1}, false},            // monotonic ;-), neither strictly increasing nor decreasing. Not trionic.
		{[]int{1, 1, 1, 1, 1}, false},      // also monotonic, except slightly longer. Not trionic.
		{[]int{1, 2, 3, 4, 5, 4, 5}, true}, // a trionic slice with long 1st phase.
		{[]int{1, 5, 4, 3, 2, 1, 3}, true}, // a trionic slice with long 2nd phase.
		{[]int{1, 2, 1, 2, 3, 4, 5}, true}, // a trionic slice with long 3rd phase.

		// a trionic slice with long 1st, 2nd, and 3rd phases.
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 3, 4, 5}, true},

		{[]int{-3, -2, -3, -2}, true},  //a trionic slice with negative numbers.
		{[]int{-3, -2, -1, -2}, false}, // not trionic, with negatives.
	}

	for i, tc := range testcases {
		actual := isTrionic(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase isTrionic#%02d (%v) failed: want %t, got %t",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
