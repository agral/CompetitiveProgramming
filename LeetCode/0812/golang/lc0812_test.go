package lc0812

import "testing"

func Test_getTriangleArea(t *testing.T) {
	testcases := []struct {
		p1       []int
		p2       []int
		p3       []int
		expected float64
	}{
		{[]int{0, 0}, []int{0, 2}, []int{2, 0}, 2.0},
		{[]int{0, 1}, []int{1, 1}, []int{0, 0}, 0.5},
	}

	for i, tc := range testcases {
		actual := getTriangleArea(tc.p1, tc.p2, tc.p3)
		if actual != tc.expected {
			t.Errorf("Testcase getTriangleArea#%02d (%v | %v | %v) failed: want %.3f, got %.3f",
				i+1, tc.p1, tc.p2, tc.p3, tc.expected, actual)
		}
	}
}

func Test_largestTriangleArea(t *testing.T) {
	testcases := []struct {
		points   [][]int
		expected float64
	}{
		{[][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}, 2.0},
		{[][]int{{1, 0}, {0, 0}, {0, 1}}, 0.5},
	}

	for i, tc := range testcases {
		actual := largestTriangleArea(tc.points)
		if actual != tc.expected {
			t.Errorf("Testcase largestTriangleArea#%02d (%v) failed: want %.3f, got %.3f",
				i+1, tc.points, tc.expected, actual)
		}
	}
}
