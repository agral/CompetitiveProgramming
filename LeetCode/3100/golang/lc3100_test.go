package lc3100

import "testing"

func Test_maxBottlesDrunk(t *testing.T) {
	testcases := []struct {
		numBottles  int
		numExchange int
		expected    int
	}{
		{13, 6, 15},
		{10, 3, 13},
	}

	for i, tc := range testcases {
		actual := maxBottlesDrunk(tc.numBottles, tc.numExchange)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (numBottles=%d, numExchange=%d) failed: want %d, got %d",
				i+1, tc.numBottles, tc.numExchange, tc.expected, actual)
		}
	}
}
