package lc3289

import (
	"slices"
	"testing"
)

func Test_getSneakyNumbers(t *testing.T) {
	testcases := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 1, 0}, []int{0, 1}},
		{[]int{0, 3, 2, 1, 3, 2}, []int{2, 3}},
		{[]int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2}, []int{4, 5}},
	}

	for i, tc := range testcases {
		actual := getSneakyNumbers(tc.input)
		slices.Sort(actual)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.input, tc.expected, actual)
		}
	}
}
