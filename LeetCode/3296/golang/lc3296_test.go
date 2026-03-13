package lc3296

import "testing"

func Test_minNumberOfSeconds(t *testing.T) {
	testcases := []struct {
		mountainHeight int
		workerTimes    []int
		expected       int64
	}{
		{4, []int{2, 1, 1}, 3},
		{10, []int{3, 2, 2, 4}, 12},
		{5, []int{1}, 15},
	}

	for i, tc := range testcases {
		actual := minNumberOfSeconds(tc.mountainHeight, tc.workerTimes)
		if actual != tc.expected {
			t.Errorf("Testcase minNumberOfSeconds#%02d (height=%d, workerTimes=%v) failed: want %d, got %d",
				i+1, tc.mountainHeight, tc.workerTimes, tc.expected, actual)
		}
	}
}
