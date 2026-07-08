package lc3756

import (
	"slices"
	"testing"
)

func Test_sumAndMultiply(t *testing.T) {
	testcases := []struct {
		s        string
		queries  [][]int
		expected []int
	}{
		{"10203004", [][]int{{0, 7}, {1, 3}, {4, 6}}, []int{12340, 4, 9}},
		{"1000", [][]int{{0, 3}, {1, 1}}, []int{1, 0}},
		{"9876543210", [][]int{{0, 9}}, []int{444444137}},
	}

	for i, tc := range testcases {
		actual := sumAndMultiply(tc.s, tc.queries)

		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase sumAndMultiply#%02d (%s) failed: want %v, got %v",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
