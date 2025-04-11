package lc2843

import "testing"

func Test_countSymmetricIntegers(t *testing.T) {
	testcases := []struct {
		low      int
		high     int
		expected int
	}{
		{1, 100, 9},     // 11, 22, 33, 44, 55, 66, 77, 88, 99
		{1200, 1230, 4}, //1203, 1212, 1221, 1230
	}

	for i, tc := range testcases {
		actual := countSymmetricIntegers(tc.low, tc.high)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d-%d) failed: want %d, got %d",
				i+1, tc.low, tc.high, tc.expected, actual)
		}
	}
}
