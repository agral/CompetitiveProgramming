package lc3633

import "testing"

func Test_earliestFinishTime(t *testing.T) {
	testcases := []struct {
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
		expected       int
	}{
		{
			[]int{2, 8},
			[]int{4, 1},
			[]int{6},
			[]int{3},
			9,
		},
		{
			[]int{5},
			[]int{3},
			[]int{1},
			[]int{10},
			14,
		},
	}

	for i, tc := range testcases {
		actual := earliestFinishTime(tc.landStartTime, tc.landDuration, tc.waterStartTime, tc.waterDuration)
		if actual != tc.expected {
			t.Errorf("Testcase #%02d failed: want %d, got %d",
				i+1, tc.expected, actual)
		}
	}
}
