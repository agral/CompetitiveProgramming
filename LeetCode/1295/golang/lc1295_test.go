package lc1295

import "testing"

func Test_findNumbers(t *testing.T) {
	testcases := []struct {
		input    []int
		expected int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}

	for i, tc := range testcases {
		actual := findNumbers(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.input, tc.expected, actual)
		}
	}
}
