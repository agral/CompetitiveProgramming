package lc1870

import "testing"

func Test_minSpeedOnTime(t *testing.T) {
	testcases := []struct {
		dist     []int
		hour     float64
		expected int
	}{
		{[]int{1, 3, 2}, 6.0, 1},
		{[]int{1, 3, 2}, 2.7, 3},
		{[]int{1, 3, 2}, 1.9, -1},
	}

	for i, tc := range testcases {
		actual := minSpeedOnTime(tc.dist, tc.hour)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %.2f) failed: want %d, got %d",
				i+1, tc.dist, tc.hour, tc.expected, actual)
		}
	}
}
