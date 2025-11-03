package lc1578

import "testing"

func Test_minCost(t *testing.T) {
	testcases := []struct {
		colors     string
		neededTime []int
		expected   int
	}{
		{"abaac", []int{1, 2, 3, 4, 5}, 3},
		{"abc", []int{1, 2, 3}, 0},
		{"aabaa", []int{1, 2, 3, 4, 1}, 2},
	}

	for i, tc := range testcases {
		actual := minCost(tc.colors, tc.neededTime)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.colors, tc.neededTime, tc.expected, actual)
		}
	}
}
