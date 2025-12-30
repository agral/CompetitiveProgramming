package lc0840

import (
	"slices"
	"testing"
)

func Test_permute(t *testing.T) {
	testcases := []struct {
		seq      []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 4, 3},
		},
		{
			[]int{1, 2, 4, 3},
			[]int{1, 3, 2, 4},
		},
		{
			[]int{1, 3, 2, 4},
			[]int{1, 3, 4, 2},
		},
		{
			[]int{1, 3, 4, 2},
			[]int{1, 4, 2, 3},
		},
		{
			[]int{1, 4, 2, 3},
			[]int{1, 4, 3, 2},
		},
		{
			[]int{1, 4, 3, 2},
			[]int{2, 1, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{1, 2, 3, 4, 5, 6, 7, 9, 8},
		},
	}

	for i, tc := range testcases {
		actual := permute(tc.seq)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.seq, tc.expected, actual)
		}
	}
}

func Test_generate3x3MagicSquares(t *testing.T) {
	squares := generate3x3MagicSquares()
	if !(8 == len(squares)) {
		t.Errorf("Wrong number of magic squares generated: want 8, got %d", len(squares))
	}

	// one magic square seen in Example #01 is:
	// 4 3 8
	// 9 5 1
	// 2 7 6
	// Check that it's present in the `squares`:
	found := false
	lookingFor := []int{4, 3, 8, 9, 1, 2, 7, 6}
	for _, magicSq := range squares {
		if slices.Equal(magicSq, lookingFor) {
			found = true
		}
	}
	if !found {
		t.Errorf("Predefined magic square %v missing from the generated array", lookingFor)
		t.Errorf("%v", squares)
	}
}

func Test_numMagicSquaresInside(t *testing.T) {
	testcases := []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{ // #01: the [0,0]->[2,2] square is magic. The [1,0]->[3,2] is not. Ans 1.
			{4, 3, 8, 4},
			{9, 5, 1, 9},
			{2, 7, 6, 2},
		}, 1},
		{[][]int{{8}}, 0}, // #02: not even 3x3. Ans 0.
		{[][]int{ // #03: not magic, but all rows/cols sum up to 15. Ans 0.
			{5, 5, 5, 5},
			{5, 5, 5, 5},
			{5, 5, 5, 5},
		}, 0},
	}

	for i, tc := range testcases {
		actual := numMagicSquaresInside(tc.grid)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.grid, tc.expected, actual)

			ms := generate3x3MagicSquares()
			t.Errorf("\n\n%v", ms)
		}
	}
}
