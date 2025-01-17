package lc2683

import "testing"

func TestDoesValidArrayExist(t *testing.T) {
	testcases := []struct {
		derived  []int
		expected bool
	}{
		{[]int{1, 1, 0}, true},
		{[]int{1, 1}, true},
		{[]int{1, 0}, false},
		{[]int{3, 3}, true},
		{[]int{3, 2}, false},
	}

	for i, tc := range testcases {
		actual := doesValidArrayExist(tc.derived)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.derived, tc.expected, actual)
		}
	}
}
