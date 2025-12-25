package lc3075

import "testing"

func Test_maximumHappinessSum(t *testing.T) {
	testcases := []struct {
		happiness []int
		k         int
		expected  int64
	}{
		{[]int{1, 2, 3}, 2, int64(4)},
		{[]int{1, 1, 1, 1}, 2, int64(1)},
		{[]int{2, 3, 4, 5}, 1, int64(5)},
	}

	for i, tc := range testcases {
		actual := maximumHappinessSum(tc.happiness, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | k=%d) failed: want %d, got %d",
				i+1, tc.happiness, tc.k, tc.expected, actual)
		}
	}
}
