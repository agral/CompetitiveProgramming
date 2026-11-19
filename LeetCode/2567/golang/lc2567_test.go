package lc2567

import "testing"

func Test_minimizeSum(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 4, 7, 8, 5}, 3},                 // {1, 4} -> {6, 6}. Then low := 6-6 = 0; high := 8-5 = 3, ans=3.
		{[]int{1, 4, 3}, 0},                       // {1, 4} -> {3, 3}. Then low := high := {3-3} = 0; ans=0.
		{[]int{1, 12, 13, 42}, 1},                 // {1, 42} -> {12, 12}. Then low := 0, high := {13-12} = 1; ans=1.
		{[]int{1, 2, 3, 4, 20, 21, 100, 101}, 20}, // {100, 101} -> {20,20}. Then low := 0,high :={21-1} = 20; ans=20.
	}

	for i, tc := range testcases {
		actual := minimizeSum(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase minimizeSum#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
