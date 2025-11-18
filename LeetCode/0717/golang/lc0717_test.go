package lc0717

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	testcases := []struct {
		bits     []int
		expected bool
	}{
		{[]int{1, 0, 0}, true},
		{[]int{1, 1, 1, 0}, false},
	}

	for i, tc := range testcases {
		actual := isOneBitCharacter(tc.bits)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %t, got %t", i+1, tc.bits, tc.expected, actual)
		}
	}
}
