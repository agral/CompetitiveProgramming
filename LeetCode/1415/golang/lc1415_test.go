package lc1415

import "testing"

func Test_getHappyString(t *testing.T) {
	testcases := []struct {
		n        int
		k        int
		expected string
	}{
		{1, 3, "c"},
		{1, 4, ""},
		{3, 9, "cab"},
	}

	for i, tc := range testcases {
		actual := getHappyString(tc.n, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase getHappyString#%02d (n=%d, k=%d) failed: want %s, got %s",
				i+1, tc.n, tc.k, tc.expected, actual)
		}
	}
}
