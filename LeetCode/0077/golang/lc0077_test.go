package lc0077

import (
	"slices"
	"testing"
)

func IsEqual(lhs [][]int, rhs [][]int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i := range lhs {
		if !slices.Equal(lhs[i], rhs[i]) {
			return false
		}
	}
	return true
}

func Test_combine(t *testing.T) {
	testcases := []struct {
		n        int
		k        int
		expected [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{1, 1, [][]int{{1}}},
	}

	for i, tc := range testcases {
		actual := combine(tc.n, tc.k)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase combine#%02d (n=%d, k=%d) failed: want %v, got %v",
				i+1, tc.n, tc.k, tc.expected, actual)
		}
	}
}
