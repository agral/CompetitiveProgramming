package lc1111

import (
	"slices"
	"testing"
)

func Test_maxDepthAfterSplit(t *testing.T) {
	testcases := []struct {
		seq      string
		expected []int
	}{
		{"(()())", []int{0, 1, 1, 1, 1, 0}},
		{"()(())()", []int{0, 0, 0, 1, 1, 0, 0, 0}}, // note: different from example 2, but also valid (!)
	}

	for i, tc := range testcases {
		actual := maxDepthAfterSplit(tc.seq)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase maxDepthAfterSplit#%02d (%s) failed: want %v, got %v",
				i+1, tc.seq, tc.expected, actual)
		}
	}
}
