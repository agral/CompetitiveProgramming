package lc3160

import (
	"slices"
	"testing"
)

func Test_queryResults(t *testing.T) {
	testcases := []struct {
		limit    int
		queries  [][]int
		expected []int
	}{
		{4, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}, []int{1, 2, 2, 3}},
		{4, [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}, []int{1, 2, 2, 3, 4}},
	}

	for i, tc := range testcases {
		actual := queryResults(tc.limit, tc.queries)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d failed: want %v, got %v", i+1, tc.expected, actual)
		}
	}
}
