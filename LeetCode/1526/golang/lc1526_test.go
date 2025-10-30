package lc1526

import "testing"

func Test_minNumberOperations(t *testing.T) {
	testcases := []struct {
		target   []int
		expected int
	}{
		{[]int{1, 2, 3, 2, 1}, 3},
		{[]int{3, 1, 1, 2}, 4},
		{[]int{3, 1, 5, 4, 2}, 7},
	}

	for i, tc := range testcases {
		actual := minNumberOperations(tc.target)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.target, tc.expected, actual)
		}
	}
}
