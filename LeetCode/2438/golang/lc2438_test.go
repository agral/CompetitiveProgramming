package lc2438

import (
	"slices"
	"testing"
)

func Test_productQueries(t *testing.T) {
	testcases := []struct {
		n        int
		queries  [][]int
		expected []int
	}{
		{15, [][]int{{0, 1}, {2, 2}, {0, 3}}, []int{2, 4, 64}},
		{2, [][]int{{0, 0}}, []int{2}},
	}

	for i, tc := range testcases {
		actual := productQueries(tc.n, tc.queries)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%d | %v) failed: want %d, got %d",
				i+1, tc.n, tc.queries, tc.expected, actual)
		}
	}
}
