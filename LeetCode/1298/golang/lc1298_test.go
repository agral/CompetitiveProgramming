package lc1298

import "testing"

func Test_maxCandies(t *testing.T) {
	testcases := []struct {
		status         []int
		candies        []int
		keys           [][]int
		containedBoxes [][]int
		initialBoxes   []int
		expected       int
	}{
		{
			/*status=*/ []int{1, 0, 1, 0},
			/*candies=*/ []int{7, 5, 4, 100},
			/*keys=*/ [][]int{{}, {}, {1}, {}},
			/*containedBoxes=*/ [][]int{{1, 2}, {3}, {}, {}},
			/*initialBoxes=*/ []int{0},
			/*expected=*/ 16,
		},
		{
			/*status=*/ []int{1, 0, 0, 0, 0, 0},
			/*candies=*/ []int{1, 1, 1, 1, 1, 1},
			/*keys=*/ [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
			/*containedBoxes=*/ [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
			/*initialBoxes=*/ []int{0},
			/*expected=*/ 6,
		},
	}

	for i, tc := range testcases {
		actual := maxCandies(tc.status, tc.candies, tc.keys, tc.containedBoxes, tc.initialBoxes)
		if actual != tc.expected {
			t.Errorf("Testcase %02d failed: want %d, got %d", i+1, tc.expected, actual)
		}
	}
}
