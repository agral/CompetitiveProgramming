package lc2169

import "testing"

func Test_countOperations(t *testing.T) {
	testcases := []struct {
		num1     int
		num2     int
		expected int
	}{
		{2, 3, 3},
		{10, 10, 1},
	}

	for i, tc := range testcases {
		actual := countOperations(tc.num1, tc.num2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d | %d) failed: want %d, got %d",
				i+1, tc.num1, tc.num2, tc.expected, actual)
		}
	}
}
