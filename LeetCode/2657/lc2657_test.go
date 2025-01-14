package lc2657

import (
	"slices"
	"testing"
)

func TestFindThePrefixCommonArray(t *testing.T) {
	testcases := []struct {
		a        []int
		b        []int
		expected []int
	}{
		{[]int{1, 3, 2, 4}, []int{3, 1, 2, 4}, []int{0, 2, 3, 4}},
		{[]int{2, 3, 1}, []int{3, 1, 2}, []int{0, 1, 3}},
	}

	for i, tc := range testcases {
		actual := findThePrefixCommonArray(tc.a, tc.b)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v | %v) failed: want %v, got %v", i+1, tc.a, tc.b, tc.expected, actual)
		}
	}
}
