package lc2558

import "testing"

func TestPickGifts(t *testing.T) {
	result1 := pickGifts([]int{25, 64, 9, 4, 100}, 4)
	expected1 := int64(29)
	if result1 != expected1 {
		t.Errorf("Example1: got %d; want %d", result1, expected1)
	}

	result2 := pickGifts([]int{1, 1, 1, 1}, 4)
	expected2 := int64(4)
	if result2 != expected2 {
		t.Errorf("Example2: got %d; want %d", result2, expected2)
	}
}
