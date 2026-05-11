package lc2553

import (
	"slices"
	"testing"
)

func Test_separateDigits(t *testing.T) {
	testcases := []struct {
		input    []int
		expected []int
	}{
		{[]int{13, 25, 83, 77}, []int{1, 3, 2, 5, 8, 3, 7, 7}},
		{[]int{7, 1, 3, 9}, []int{7, 1, 3, 9}},
		{[]int{10, 1000, 1000000}, []int{1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}},
	}

	for i, tc := range testcases {
		actual := separateDigits(tc.input)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase separatDigits#%02d (%v) failed: want %v, got %v",
				i+1, tc.input, tc.expected, actual)
		}
	}
}
