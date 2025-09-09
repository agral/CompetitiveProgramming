package lc2327

import "testing"

func Test_peopleAwareOfSecret(t *testing.T) {
	testcases := []struct {
		n        int
		delay    int
		forget   int
		expected int
	}{
		{6, 2, 4, 5},
		{4, 1, 3, 6},
	}

	for i, tc := range testcases {
		actual := peopleAwareOfSecret(tc.n, tc.delay, tc.forget)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.n, tc.expected, actual)
		}
	}
}
