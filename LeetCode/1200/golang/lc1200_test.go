package lc1200

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

func Test_minimumAbsDifference(t *testing.T) {
	testcases := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{4, 2, 1, 3}, [][]int{{1, 2}, {2, 3}, {3, 4}}},
		{[]int{1, 3, 6, 10, 15}, [][]int{{1, 3}}},
		{[]int{3, 8, -10, 23, 19, -4, -14, 27}, [][]int{{-14, -10}, {19, 23}, {23, 27}}},
	}

	for i, tc := range testcases {
		actual := minimumAbsDifference(tc.input)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase minimumAbsDifference#%02d (%v) failed: want %v, got %v",
				i+1, tc.input, tc.expected, actual)
		}
	}
}
