package lc1689

import "testing"

func Test_minPartitions(t *testing.T) {
	testcases := []struct {
		n        string
		expected int
	}{
		{"32", 3},
		{"82734", 8},
		{"27346209830709182346", 9},
		{"45", 5},
		{"1000000001", 1},
	}

	for i, tc := range testcases {
		actual := minPartitions(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase minPartitions#%02d (%s) failed: want %d, got %d",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
