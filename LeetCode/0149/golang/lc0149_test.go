package lc0149

import "testing"

func Test_gcd(t *testing.T) {
	testcases := []struct {
		x        int
		y        int
		expected int
	}{
		{10, 100, 10},
		{100, 10, 10},
		{101, 101, 101},
		{9, 10, 1},
		{19, 1024, 1},
		{15, 21, 3},
	}

	for i, tc := range testcases {
		actual := gcd(tc.x, tc.y)
		if actual != tc.expected {
			t.Errorf("Testcase gcd#%02d (%d, %d) failed: want %d, got %d",
				i+1, tc.x, tc.y, tc.expected, actual)
		}
	}
}

func Test_maxPoints(t *testing.T) {
	testcases := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
		{[][]int{{3, 3}, {1, 4}, {1, 1}, {2, 1}, {2, 2}}, 3},
	}

	for i, tc := range testcases {
		actual := maxPoints(tc.points)
		if actual != tc.expected {
			t.Errorf("Testcase maxPoints#%02d (%v) failed: want %d, got %d",
				i+1, tc.points, tc.expected, actual)
		}
	}
}
