package lc2466

import "testing"

func TestCountGoodStrings(t *testing.T) {
	data := [][]int{ // 0: low, 1: high, 2: zero, 3: one, 4: expected
		{3, 3, 1, 1, 8},
		{2, 3, 1, 2, 5},
	}

	for i, d := range data {
		result := countGoodStrings(d[0], d[1], d[2], d[3])
		if result != d[4] {
			t.Errorf("Example %02d: got %v, want %v", i+1, result, d[4])
		}
	}
}
