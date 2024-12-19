package lc0026

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	inputs := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{-5, -3, -1, 1, 3, 5},
		{1},
		{1, 1, 1, 1, 1},
	}
	expected := []int{
		2,
		5,
		6,
		1,
		1,
	}

	for i := range inputs {
		result := removeDuplicates(inputs[i])
		if result != expected[i] {
			t.Errorf("Example%02d: got %v, want %v", i+1, result, expected[i])
		}
	}
}
