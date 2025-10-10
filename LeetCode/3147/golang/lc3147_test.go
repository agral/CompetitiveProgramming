package lc3147

import "testing"

func Test_maximumEnergy(t *testing.T) {
	testcases := []struct {
		energy   []int
		k        int
		expected int
	}{
		{[]int{5, 2, -10, -5, 1}, 3, 3},
		{[]int{-2, -3, -1}, 2, -1},
	}

	for i, tc := range testcases {
		actual := maximumEnergy(tc.energy, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.energy, tc.k, tc.expected, actual)
		}
	}
}
