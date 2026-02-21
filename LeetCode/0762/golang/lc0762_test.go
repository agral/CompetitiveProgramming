package lc0762

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	testcases := []struct {
		left     int
		right    int
		expected int
	}{
		{6, 10, 4},  // 6, 7, 9, 10
		{10, 15, 5}, // 10-14 have prime binary repr; only 15 (1111) does not.
	}

	for i, tc := range testcases {
		actual := countPrimeSetBits(tc.left, tc.right)
		if actual != tc.expected {
			t.Errorf("Testcase countPrimeSetBits#%02d (%d-%d) failed: want %d, got %d",
				i+1, tc.left, tc.right, tc.expected, actual)
		}
	}
}
