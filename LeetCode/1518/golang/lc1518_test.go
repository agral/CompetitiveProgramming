package lc1518

import "testing"

func Test_numWaterBottles(t *testing.T) {
	testcases := []struct {
		numBottles  int
		numExchange int
		expected    int
	}{
		{9, 3, 13},
		{15, 4, 19},
	}

	for i, tc := range testcases {
		actual := numWaterBottles(tc.numBottles, tc.numExchange)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (numBottles=%d, numExchange=%d) failed: want %d, got %d",
				i+1, tc.numBottles, tc.numExchange, tc.expected, actual)
		}
	}
}
