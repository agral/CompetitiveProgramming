package lc2011

import "testing"

func Test_finalValueAfterOperations(t *testing.T) {
	testcases := []struct {
		operations []string
		expected   int
	}{
		{[]string{"--X", "X++", "X++"}, 1},
		{[]string{"++X", "++X", "X++"}, 3},
		{[]string{"X++", "++X", "--X", "X--"}, 0},
	}

	for i, tc := range testcases {
		actual := finalValueAfterOperations(tc.operations)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.operations, tc.expected, actual)
		}
	}
}
