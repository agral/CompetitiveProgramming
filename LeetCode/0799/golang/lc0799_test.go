package lc0799

import "testing"

func Test_champagneTower(t *testing.T) {
	testcases := []struct {
		poured      int
		query_row   int
		query_glass int
		expected    float64
	}{
		{1, 1, 1, 0.0},
		{2, 1, 1, 0.5},
		{100000009, 33, 17, 1.0},
	}

	for i, tc := range testcases {
		actual := champagneTower(tc.poured, tc.query_row, tc.query_glass)
		if actual != tc.expected {
			t.Errorf("Testcase champagneTower#%02d (poured=%d, row=%d, glass=%d) failed: want %.5f, got %.5f",
				i+1, tc.poured, tc.query_row, tc.query_glass, tc.expected, actual)
		}
	}
}
