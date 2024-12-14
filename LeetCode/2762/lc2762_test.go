package lc2762

import "testing"

func TestContinuousSubarrays(t *testing.T) {
	result1 := continuousSubarrays([]int{5, 4, 2, 4})
	expected1 := int64(8)
	if result1 != expected1 {
		t.Errorf("Example1: got %d, want %d", result1, expected1)
	}

	result2 := continuousSubarrays([]int{1, 2, 3})
	expected2 := int64(6)
	if result2 != expected2 {
		t.Errorf("Example1: got %d, want %d", result2, expected2)
	}
}
