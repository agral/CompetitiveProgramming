package lc2433

import (
	"slices"
	"testing"
)

func Test_findArray(t *testing.T) {
	testcases := []struct {
		pref     []int
		expected []int
	}{
		{[]int{5, 2, 0, 3, 1}, []int{5, 7, 2, 3, 2}},
		{[]int{13}, []int{13}},
	}

	for i, tc := range testcases {
		actual := findArray(tc.pref)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase findArray#%02d (%v) failed: want %v, got %v",
				i+1, tc.pref, tc.expected, actual)
		}
	}
}
