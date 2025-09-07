package lc1304

import "testing"

func SumsToZero(s []int) bool {
	ans := 0
	for _, num := range s {
		ans += num
	}
	return ans == 0
}

func Test_sumZero(t *testing.T) {
	testcases := []struct {
		n int
	}{
		{5},
		{4},
		{3},
		{2},
		{1},
	}

	for i, tc := range testcases {
		actual := sumZero(tc.n)
		if !SumsToZero(actual) {
			t.Errorf("Testcase %02d (%v) failed: want a slice that sums to zero, got %d", i+1, tc.n, actual)
		}
	}
}
