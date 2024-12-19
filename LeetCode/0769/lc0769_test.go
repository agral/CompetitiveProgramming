package lc0769

import "testing"

func TestMaxChunksToSorted(t *testing.T) {
	inputs := [][]int{
		{4, 3, 2, 1, 0},
		{1, 0, 2, 3, 4},
	}
	expected := []int{
		1,
		4,
	}

	for i := range inputs {
		result := maxChunksToSorted(inputs[i])
		if result != expected[i] {
			t.Errorf("Example%02d: got %v, want %v", i+1, result, expected[i])
		}
	}
}
