package lc2946

import "testing"

func Test_areSimilar(t *testing.T) {
	testcases := []struct {
		mat      [][]int
		k        int
		expected bool
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 4, false,
		},
		{
			[][]int{
				{1, 2, 1, 2},
				{5, 5, 5, 5},
				{6, 3, 6, 3},
			}, 2, true,
		},
	}

	for i, tc := range testcases {
		actual := areSimilar(tc.mat, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase areSimilar#%02d (mat=%v, k=%d) failed: want %t, got %t",
				i+1, tc.mat, tc.k, tc.expected, actual)
		}
	}
}
