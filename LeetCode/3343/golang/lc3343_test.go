package lc3343

import (
	"slices"
	"testing"
)

func Test_countBalancedPermutations(t *testing.T) {
	testcases := []struct {
		num      string
		expected int
	}{
		{"123", 2},
		{"112", 1},
		{"12345", 0},
	}

	for i, tc := range testcases {
		actual := countBalancedPermutations(tc.num)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %d, got %d", i+1, tc.num, tc.expected, actual)
		}
	}
}

func Test_factorial(t *testing.T) {
	testcases := []struct {
		input    int
		expected int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{10, 3628800},
	}
	for _, tc := range testcases {
		actual := factorial(tc.input)
		if actual != tc.expected {
			t.Errorf("factorial(%d) failed: want %d, got %d", tc.input, tc.expected, actual)
		}
	}
}

func Test_getDigits(t *testing.T) {
	testcases := []struct {
		input    string
		expected []int
	}{
		{"1", []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"123", []int{0, 1, 1, 1, 0, 0, 0, 0, 0, 0}},
		{"123122111", []int{0, 5, 3, 1, 0, 0, 0, 0, 0, 0}},
	}

	for _, tc := range testcases {
		actual := getDigits(tc.input)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("getDigits(%s) failed: want %v, got %v", tc.input, tc.expected, actual)
		}
	}
}
