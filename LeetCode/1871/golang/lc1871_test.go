package lc1871

import "testing"

func Test_canReach(t *testing.T) {
	testcases := []struct {
		s        string
		minJump  int
		maxJump  int
		expected bool
	}{
		{"011010", 2, 3, true},
		{"01101110", 2, 3, false},
	}

	for i, tc := range testcases {
		actual := canReach(tc.s, tc.minJump, tc.maxJump)
		if actual != tc.expected {
			t.Errorf("Testcase canReach#%02d (%s, minJump=%d, maxJump=%d) failed: want %t, got %t",
				i+1, tc.s, tc.minJump, tc.maxJump, tc.expected, actual)
		}
	}
}
