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
