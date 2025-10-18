package lc3304

import "testing"

func Test_kthCharacter(t *testing.T) {
	testcases := []struct {
		k        int
		expected byte
	}{
		{5, byte('b')},
		{10, byte('c')},
	}

	for i, tc := range testcases {
		actual := kthCharacter(tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d) failed: want %v, got %v", i+1, tc.k, tc.expected, actual)
		}
	}
}
