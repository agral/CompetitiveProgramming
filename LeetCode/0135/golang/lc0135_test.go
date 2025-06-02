package lc0135

import "testing"

func Test_candy(t *testing.T) {
	testcases := []struct {
		ratings  []int
		expected int
	}{
		{[]int{1, 0, 2}, 5}, // allocate: [2, 1, 2]
		{[]int{1, 2, 2}, 4}, // allocate: [1, 2, 1]
	}

	for i, tc := range testcases {
		actual := candy(tc.ratings)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.ratings, tc.expected, actual)
		}
	}
}
