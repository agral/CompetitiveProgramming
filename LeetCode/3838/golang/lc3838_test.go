package lc3838

import "testing"

func Test_mapWordWeights(t *testing.T) {
	testcases := []struct {
		words    []string
		weights  []int
		expected string
	}{
		{
			[]string{"abcd", "def", "xyz"},
			[]int{5, 3, 12, 14, 1, 2, 3, 2, 10, 6, 6, 9, 7, 8, 7, 10, 8, 9, 6, 9, 9, 8, 3, 7, 7, 2},
			"rij",
		},
		{
			[]string{"a", "b", "c"},
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			"yyy",
		},
		{
			[]string{"abcd"},
			[]int{7, 5, 3, 4, 3, 5, 4, 9, 4, 2, 2, 7, 10, 2, 5, 10, 6, 1, 2, 2, 4, 1, 3, 4, 4, 5},
			"g",
		},
	}

	for i, tc := range testcases {
		actual := mapWordWeights(tc.words, tc.weights)
		if actual != tc.expected {
			t.Errorf("Testcase mapWordWeights#%02d (%v) failed: want %v, got %v",
				i+1, tc.words, tc.expected, actual)
		}
	}
}
