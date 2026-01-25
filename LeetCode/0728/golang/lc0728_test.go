package lc0728

import (
	"slices"
	"testing"
)

func Test_isSelfDividing(t *testing.T) {
	testcases := []struct {
		num      int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{9, true},
		{10, false},
		{11, true},
		{12, true},
		{13, false},
		{14, false},
		{15, true},
		{20, false},
		{21, false},
		{22, true},
		{127, false},
		{128, true},
	}

	for i, tc := range testcases {
		actual := isSelfDividing(tc.num)
		if actual != tc.expected {
			t.Errorf("Testcase isSelfDividing#%02d (%v) failed: want %t, got %t",
				i+1, tc.num, tc.expected, actual)
		}
	}
}

func Test_selfDividingNumbers(t *testing.T) {
	testcases := []struct {
		left     int
		right    int
		expected []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{47, 85, []int{48, 55, 66, 77}},
	}

	for i, tc := range testcases {
		actual := selfDividingNumbers(tc.left, tc.right)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase selfDividingNumbers#%02d ([%d-%d]) failed: want %v, got %v",
				i+1, tc.left, tc.right, tc.expected, actual)
		}
	}
}
