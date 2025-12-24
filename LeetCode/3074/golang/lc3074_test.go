package lc3074

import "testing"

func Test_minimumBoxes(t *testing.T) {
	testcases := []struct {
		apples   []int
		boxes    []int
		expected int
	}{
		{[]int{1, 3, 2}, []int{4, 3, 1, 5, 2}, 2},
		{[]int{5, 5, 5}, []int{2, 4, 2, 7}, 4},
	}

	for i, tc := range testcases {
		actual := minimumBoxes(tc.apples, tc.boxes)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (apples=%v | boxes=%v) failed: want %d, got %d",
				i+1, tc.apples, tc.boxes, tc.expected, actual)
		}
	}
}
