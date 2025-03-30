package lc0763

import (
	"slices"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	testcases := []struct {
		s        string
		expected []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"eccbbbbdec", []int{10}},
	}

	for i, tc := range testcases {
		actual := partitionLabels(tc.s)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.s, tc.expected, actual)
		}
	}
}
