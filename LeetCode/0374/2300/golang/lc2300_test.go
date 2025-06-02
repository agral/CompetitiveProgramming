package lc2300

import (
	"slices"
	"testing"
)

func Test_successfulPairs(t *testing.T) {
	testcases := []struct {
		spells   []int
		potions  []int
		success  int64
		expected []int
	}{
		{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7, []int{4, 0, 3}},
		{[]int{3, 1, 2}, []int{8, 5, 8}, 16, []int{2, 0, 2}},
	}

	for i, tc := range testcases {
		actual := successfulPairs(tc.spells, tc.potions, tc.success)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d failed: want %v, got %v", i+1, tc.expected, actual)
		}
	}
}
