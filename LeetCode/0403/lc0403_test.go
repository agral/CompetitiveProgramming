package lc0403

import "testing"

func TestCanCross(t *testing.T) {
	testingData := []struct {
		stones   []int
		expected bool
	}{
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
		{[]int{0, 1}, true},
		{[]int{0, 1, 2}, true},
		{[]int{0, 1, 3}, true},
		{[]int{0, 1, 3, 4}, true},
		{[]int{0, 1, 3, 5}, true},
		{[]int{0, 1, 3, 6}, true},
		{[]int{0, 1, 3, 7}, false},
	}

	for i, d := range testingData {
		actual := canCross(d.stones)
		if actual != d.expected {
			t.Errorf("Testcase %02d (%v): want %v, got %v", i, d.stones, d.expected, actual)
		}
	}
}
