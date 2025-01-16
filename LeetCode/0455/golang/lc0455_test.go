package lc0455

import "testing"

func TestFindContentChildren(t *testing.T) {
	testcases := []struct {
		g        []int
		s        []int
		expected int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}

	for i, tc := range testcases {
		actual := findContentChildren(tc.g, tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.g, tc.s, tc.expected, actual)
		}
	}
}
