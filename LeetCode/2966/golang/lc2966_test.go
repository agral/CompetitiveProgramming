package lc2966

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

func Test_divideArray(t *testing.T) {
	testcases := []struct {
		nums     []int
		k        int
		expected [][]int
	}{
		{
			/*nums=*/ []int{1, 3, 4, 8, 7, 9, 3, 5, 1},
			/*k=*/ 2,
			/*expected=*/ [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}},
		},
		{
			/*nums=*/ []int{1, 3, 3, 2, 7, 3},
			/*k=*/ 2,
			/*expected=*/ [][]int{},
		},
		{
			/*nums=*/ []int{4, 2, 9, 8, 2, 12, 7, 12, 10, 5, 8, 5, 5, 7, 9, 2, 5, 11},
			/*k=*/ 14,
			// /*expected=*/ [][]int{{2, 2, 12}, {4, 8, 5}, {5, 9, 7}, {7, 8, 5}, {5, 9, 10}, {11, 12, 2}},
			/*expected=*/ [][]int{{2, 2, 2}, {4, 5, 5}, {5, 5, 7}, {7, 8, 8}, {9, 9, 10}, {11, 12, 12}},
		},
	}

	for i, tc := range testcases {
		actual := divideArray(tc.nums, tc.k)
		if !IsEqual(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v | %d) failed: want %v, got %v",
				i+1, tc.nums, tc.k, tc.expected, actual)
		}
	}
}
