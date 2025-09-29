package lc1039

import "testing"

func Test_minScoreTriangulation(t *testing.T) {
	testcases := []struct {
		values   []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{3, 7, 4, 5}, 144},
		{[]int{1, 3, 1, 4, 1, 5}, 13},
	}

	for i, tc := range testcases {
		actual := minScoreTriangulation(tc.values)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.values, tc.expected, actual)
		}
	}
}
