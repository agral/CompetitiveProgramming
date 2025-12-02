package lc3623

import "testing"

func Test_countTrapezoids(t *testing.T) {
	testcases := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}}, 3},
		{[][]int{{0, 0}, {0, 1}, {1, 0}, {2, 1}}, 1},

		// all points with same y-coordinate:
		{[][]int{{0, 9}, {1, 9}, {2, 9}, {3, 9}, {4, 9}}, 0},

		// all points with different y-coordinate:
		{[][]int{{9, 0}, {9, 1}, {9, 2}, {9, 3}, {9, 4}}, 0},

		// WA #01:
		{[][]int{{-73, -72}, {-1, -56}, {-92, 30}, {-57, -89}, {-19, -89}, {-35, 30}, {64, -72}}, 3},
	}

	for i, tc := range testcases {
		actual := countTrapezoids(tc.points)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.points, tc.expected, actual)
		}
	}
}
