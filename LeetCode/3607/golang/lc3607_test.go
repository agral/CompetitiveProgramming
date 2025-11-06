package lc3607

import (
	"slices"
	"testing"
)

func Test_processQueries(t *testing.T) {
	testcases := []struct {
		c           int
		connections [][]int
		queries     [][]int
		expected    []int
	}{
		{
			/* c */ 5,
			/* connections */ [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			/* queries */ [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}},
			/* expected */ []int{3, 2, 3},
		},
		{
			/* c */ 3,
			/* connections */ [][]int{{}},
			/* queries */ [][]int{{1, 1}, {2, 1}, {1, 1}},
			/* expected */ []int{1, -1},
		},
	}

	for i, tc := range testcases {
		actual := processQueries(tc.c, tc.connections, tc.queries)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
