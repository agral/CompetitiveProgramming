package lc0051

import (
	"slices"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	testcases := []struct {
		n        int
		expected [][]string
	}{
		{1, [][]string{{"Q"}}},
		{2, [][]string{}},
		{3, [][]string{}},
		{4, [][]string{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		}},
	}

	isSameSetup := func(expected [][]string, tested [][]string) bool {
		if len(expected) != len(tested) {
			return false
		}
		for i := 0; i < len(expected); i++ {
			if !slices.Equal(expected[i], tested[i]) {
				return false
			}
		}
		return true
	}

	for i, tc := range testcases {
		actual := solveNQueens(tc.n)
		if !isSameSetup(tc.expected, actual) {
			t.Errorf("Testcase %02d (%v) failed: want %v, got %v", i+1, tc.n, tc.expected, actual)
		}
	}
}
