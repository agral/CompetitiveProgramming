package lc3494

import "testing"

func Test_minTime(t *testing.T) {
	testcases := []struct {
		skill    []int
		mana     []int
		expected int64
	}{
		{[]int{1, 5, 2, 4}, []int{5, 1, 4, 2}, 110},
		{[]int{1, 1, 1}, []int{1, 1, 1}, 5},
		{[]int{1, 2, 3, 4}, []int{1, 2}, 21},
	}

	for i, tc := range testcases {
		actual := minTime(tc.skill, tc.mana)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.skill, tc.mana, tc.expected, actual)
		}
	}
}
