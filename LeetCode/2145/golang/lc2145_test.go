package lc2145

import "testing"

func Test_numberOfArrays(t *testing.T) {
	testcases := []struct {
		differences []int
		lower       int
		upper       int
		expected    int
	}{
		{[]int{1, -3, 4}, 1, 6, 2},
		{[]int{3, -4, 5, 1, -2}, -4, 5, 4},
		{[]int{4, -7, 2}, 3, 6, 0},
		{[]int{1, 1}, 0, 2, 1}, // [0, 1, 2]
	}

	for i, tc := range testcases {
		actual := numberOfArrays(tc.differences, tc.lower, tc.upper)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.differences, tc.expected, actual)
		}
	}
}
