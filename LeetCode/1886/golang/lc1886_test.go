package lc1886

import "testing"

func Test_findRotation(t *testing.T) {
	testcases := []struct {
		mat      [][]int
		target   [][]int
		expected bool
	}{
		{
			[][]int{
				{0, 1},
				{1, 0},
			},
			[][]int{
				{1, 0},
				{0, 1},
			}, true,
		},
		{
			[][]int{
				{0, 1},
				{1, 1},
			},
			[][]int{
				{1, 0},
				{0, 1},
			}, false,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			[][]int{
				{1, 1, 1},
				{0, 1, 0},
				{0, 0, 0},
			}, true,
		},
	}

	for i, tc := range testcases {
		actual := findRotation(tc.mat, tc.target)
		if actual != tc.expected {
			t.Errorf("Testcase findRotation#%02d:mat=\n%v\ntarget=\n%v\nfailed: want %t, got %t",
				i+1, tc.mat, tc.target, tc.expected, actual)
		}
	}
}
