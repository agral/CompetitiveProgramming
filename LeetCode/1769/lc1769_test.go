package lc1769

import (
	"slices"
	"testing"
)

func TestMinOperations(t *testing.T) {
	testingData := []struct {
		boxes    string
		expected []int
	}{
		{"110", []int{1, 1, 3}},
		{"001011", []int{11, 8, 5, 4, 3, 4}},
	}

	for i, d := range testingData {
		actual := minOperations(d.boxes)
		if !slices.Equal(actual, d.expected) {
			t.Errorf("Testcase %02d (%s) failed: want %v, got %v", i, d.boxes, d.expected, actual)
		}
	}
}
