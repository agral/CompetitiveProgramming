package lc1979

import "testing"

func Test_findGCD(t *testing.T) {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 5, 6, 9, 10}, 2}, // gcd(2, 10) == 2
		{[]int{7, 5, 6, 8, 3}, 1},  // gcd(3, 8) == 1
		{[]int{3, 3}, 3},           // gcd(3, 3) == 3
	}

	for i, tc := range testcases {
		actual := findGCD(tc.nums)
		if actual != tc.expected {
			t.Errorf("Testcase findGCD#%02d (%v) failed: want %d, got %d",
				i+1, tc.nums, tc.expected, actual)
		}
	}
}
