package lc2094

import (
	"slices"
	"testing"
)

func Test_findEvenNumbers(t *testing.T) {
	testcases := []struct {
		digits   []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{132, 312}},
		{[]int{2, 1, 3, 0}, []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}},
		{[]int{2, 2, 8, 8, 2}, []int{222, 228, 282, 288, 822, 828, 882}},
		{[]int{3, 7, 5}, []int{}},
	}

	for i, tc := range testcases {
		actual := findEvenNumbers(tc.digits)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v",
				i, tc.digits, tc.expected, actual)
		}
	}
}
