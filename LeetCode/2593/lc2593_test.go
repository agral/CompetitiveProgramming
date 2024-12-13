package lc2593

import "testing"

func TestFindScore(t *testing.T) {
	result1 := findScore([]int{2, 1, 3, 4, 5, 2})
	expected1 := int64(7)
	if result1 != expected1 {
		t.Errorf("Example1: got %d; want %d", result1, expected1)
	}
	result2 := findScore([]int{2, 3, 5, 1, 3, 2})
	expected2 := int64(5)
	if result2 != expected2 {
		t.Errorf("Example2: got %d; want %d", result2, expected2)
	}
	result_wa1 := findScore([]int{2, 5, 6, 6, 10})
	expected_wa1 := int64(18)
	if result_wa1 != expected_wa1 {
		t.Errorf("WA1: got %d; want %d", result_wa1, expected_wa1)
	}
}
