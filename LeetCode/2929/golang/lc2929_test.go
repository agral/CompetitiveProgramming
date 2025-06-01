package lc2929

import "testing"

func Test_distributeCandies(t *testing.T) {
	testcases := []struct {
		n        int
		limit    int
		expected int64
	}{
		{5, 2, int64(3)},
		{3, 3, int64(10)},
	}

	for i, tc := range testcases {
		actual := distributeCandies(tc.n, tc.limit)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d | %d) failed: want %d, got %d",
				i+1, tc.n, tc.limit, tc.expected, actual)
		}
	}
}
