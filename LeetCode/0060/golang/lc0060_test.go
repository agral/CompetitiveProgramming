package lc0060

import "testing"

func Test_getPermutation(t *testing.T) {
	testcases := []struct {
		n        int
		k        int
		expected string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
		{3, 1, "123"},
	}

	for i, tc := range testcases {
		actual := getPermutation(tc.n, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase getPermutation#%02d (n=%d, k=%d) failed: want %s, got %s",
				i+1, tc.n, tc.k, tc.expected, actual)
		}
	}
}
