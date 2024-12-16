package lc3264

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

func TestGetFinalState(t *testing.T) {
	result1 := getFinalState([]int{2, 1, 3, 5, 6}, 5, 2)
	expected1 := []int{8, 4, 6, 5, 6}
	if !IsSlicesEqual(result1, expected1) {
		t.Errorf("Example1: got %v, want %v", result1, expected1)
	}

	result2 := getFinalState([]int{1, 2}, 3, 4)
	expected2 := []int{16, 8}
	if !IsSlicesEqual(result2, expected2) {
		t.Errorf("Example1: got %v, want %v", result2, expected2)
	}
}
