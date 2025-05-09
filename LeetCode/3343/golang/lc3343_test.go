package lc3343

import "testing"

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
