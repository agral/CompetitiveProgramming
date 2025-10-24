package lc2048

import "testing"

func Test_nextBeautifulNumber(t *testing.T) {
	testcases := []struct {
		num      int
		expected int
	}{
		{1, 22},
		{1000, 1333},
		{3000, 3133},
	}

	for i, tc := range testcases {
		actual := nextBeautifulNumber(tc.num)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.num, tc.expected, actual)
		}
	}
}
