package lc2483

import "testing"

func Test_bestClosingTime(t *testing.T) {
	testcases := []struct {
		customers string
		expected  int
	}{
		{"YYNY", 2},
		{"NNNNN", 0},
		{"YYYY", 4},
		{"NYYYNNNYNN", 4},
	}

	for i, tc := range testcases {
		actual := bestClosingTime(tc.customers)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %d, got %d", i+1, tc.customers, tc.expected, actual)
		}
	}
}
