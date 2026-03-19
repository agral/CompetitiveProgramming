package lc3212

import "testing"

func Test_numberOfSubmatrices(t *testing.T) {
	testcases := []struct {
		grid     [][]byte
		expected int
	}{
		{[][]byte{
			{'X', 'Y', '.'},
			{'Y', '.', '.'},
		}, 3},

		{[][]byte{
			{'X', 'X'},
			{'X', 'Y'},
		}, 0},

		{[][]byte{
			{'.', '.'},
			{'.', '.'},
		}, 0},
	}

	for i, tc := range testcases {
		actual := numberOfSubmatrices(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase numberOfSubmatrices#%02d (%v) failed: want %d, got %d",
				i+1, tc.grid, tc.expected, actual)
		}
	}
}
