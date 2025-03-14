package lc2226

import "testing"

func Test_maximumCandies(t *testing.T) {
	testcases := []struct {
		candies  []int
		k        int64
		expected int
	}{
		{[]int{5, 8, 6}, 3, 5},
		{[]int{2, 5}, 11, 0},
	}

	for i, tc := range testcases {
		actual := maximumCandies(tc.candies, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v; %d) failed: want %d, got %d",
				i+1, tc.candies, tc.k, tc.expected, actual)
		}
	}
}
