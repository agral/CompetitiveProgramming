package lc1475

import "testing"

// Checks whether slices lhs and rhs contain the same elements.
func IsSlicesEqual(lhs, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for idx, val := range lhs {
		if val != rhs[idx] {
			return false
		}
	}
	return true
}

func TestFinalPrices(t *testing.T) {
	inputs := [][]int{
		[]int{8, 4, 6, 2, 3},
		[]int{1, 2, 3, 4, 5},
		[]int{10, 1, 1, 6},
	}
	expected := [][]int{
		[]int{4, 2, 4, 2, 3},
		[]int{1, 2, 3, 4, 5},
		[]int{9, 0, 1, 6},
	}

	for i := range len(inputs) {
		result := finalPrices(inputs[i])
		if !IsSlicesEqual(result, expected[i]) {
			t.Errorf("Example1: got %v, want %v", result, expected[i])
		}
	}
}
