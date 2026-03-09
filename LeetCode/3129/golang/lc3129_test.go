package lc3129

import "testing"

func Test_numberOfStableArrays(t *testing.T) {
	testcases := []struct {
		zero     int
		one      int
		limit    int
		expected int
	}{
		{1, 1, 2, 2},  // [1, 0], and [0, 1] are the only such stable binary arrays.
		{1, 2, 1, 1},  // [1, 0, 1] - the only one.
		{3, 3, 2, 14}, // [1, 1, 1, 0, 0, 0] and [0, 0, 0, 1, 1, 1] are not stable; the other 14 are.
	}

	for i, tc := range testcases {
		actual := numberOfStableArrays(tc.zero, tc.one, tc.limit)
		if actual != tc.expected {
			t.Errorf("Testcase numberOfStableArrays#%02d (zero=%d, one=%d, limit=%d) failed: want %d, got %d",
				i+1, tc.zero, tc.one, tc.limit, tc.expected, actual)
		}
	}
}
