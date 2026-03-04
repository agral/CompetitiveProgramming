package lc1582

import "testing"

func Test_numSpecial(t *testing.T) {
	testcases := []struct {
		mat      [][]int
		expected int
	}{
		{[][]int{
			{1, 0, 0},
			{0, 0, 1},
			{1, 0, 0}}, 1},
		{[][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1}}, 3},
		{[][]int{
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0}}, 0},
		{[][]int{
			{1, 0, 0, 0, 1},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 0, 0}}, 1},
		{[][]int{
			{0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 1, 0, 0, 1},
			{0, 0, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0, 0},
			{0, 0, 1, 1, 0, 0, 0, 0}}, 1},
	}

	for i, tc := range testcases {
		actual := numSpecial(tc.mat)
		if actual != tc.expected {
			t.Errorf("Testcase numSpecial#%02d (%v) failed: want %d, got %d",
				i+1, tc.mat, tc.expected, actual)
		}
	}
}
