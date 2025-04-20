package lc0781

import "testing"

func Test_numRabbits(t *testing.T) {
	testcases := []struct {
		answers  []int
		expected int
	}{
		{[]int{1, 1, 2}, 5},
		{[]int{10, 10, 10}, 11},
	}

	for i, tc := range testcases {
		actual := numRabbits(tc.answers)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.answers, tc.expected, actual)
		}
	}
}
