package lc3653

import "testing"

func Test_xorAfterQueries(t *testing.T) {
	testcases := []struct {
		nums     []int
		queries  [][]int
		expected int
	}{
		{
			[]int{1, 1, 1},
			[][]int{{0, 2, 1, 4}},
			4,
		},
		{
			[]int{2, 3, 1, 5, 4},
			[][]int{{1, 4, 2, 3}, {0, 2, 1, 2}},
			31,
		},
		{
			[]int{780},
			[][]int{{0, 0, 1, 13}, {0, 0, 1, 17}, {0, 0, 1, 9}, {0, 0, 1, 18}, {0, 0, 1, 16},
				{0, 0, 1, 6}, {0, 0, 1, 4}, {0, 0, 1, 11}, {0, 0, 1, 7}, {0, 0, 1, 18},
				{0, 0, 1, 8}, {0, 0, 1, 15}, {0, 0, 1, 12}},
			523618060,
		},
	}

	for i, tc := range testcases {
		actual := xorAfterQueries(tc.nums, tc.queries)
		if actual != tc.expected {
			t.Errorf("Testcase xorAfterQueries#%02d (nums=%v, queries=%v) failed: want %d, got %d",
				i+1, tc.nums, tc.queries, tc.expected, actual)
		}
	}
}
