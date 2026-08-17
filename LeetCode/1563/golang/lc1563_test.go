package lc1563

import "testing"

func Test_stoneGameV(t *testing.T) {
	testcases := []struct {
		stoneValue []int
		expected   int
	}{
		{[]int{6, 2, 3, 4, 5, 5}, 18},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 28},
		{[]int{4}, 0},
		{[]int{71826576}, 0},
	}

	for i, tc := range testcases {
		actual := stoneGameV(tc.stoneValue)
		if actual != tc.expected {
			t.Errorf("Testcase stoneGameV#%02d (%v) failed: want %d, got %d",
				i+1, tc.stoneValue, tc.expected, actual)
		}
	}
}
