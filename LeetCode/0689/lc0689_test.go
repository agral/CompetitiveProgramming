package lc0689

import (
	"slices"
	"testing"
)

func TestMaxSumOfThreeSubarrays(t *testing.T) {
	inputs := [][]int{
		{1, 2, 1, 2, 6, 7, 5, 1},
		{1, 2, 1, 2, 1, 2, 1, 2, 1},
	}
	ks := []int{
		2,
		2,
	}
	expected := [][]int{
		{0, 3, 5},
		{0, 2, 4},
	}

	for i := range inputs {
		result := maxSumOfThreeSubarrays(inputs[i], ks[i])
		if !slices.Equal(result, expected[i]) {
			t.Errorf("Example%02d: got %v, want %v", i+1, result, expected[i])
		}
	}

}
