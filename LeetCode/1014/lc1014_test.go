package lc1014

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	inputs := [][]int{
		{8, 1, 5, 2, 6},
		{1, 2},
	}
	expected := []int{
		11,
		2,
	}

	for i := range inputs {
		result := maxScoreSightseeingPair(inputs[i])
		if result != expected[i] {
			t.Errorf("Example%02d: got %v, want %v", i+1, result, expected)
		}
	}
}
