package lc3498

import "testing"

func Test_reverseDegree(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{"abc", 148},
		{"zaza", 160},
	}

	for i, tc := range testcases {
		actual := reverseDegree(tc.s)
		if actual != tc.expected {
			t.Errorf("Testcase reverseDegree#%02d (%v) failed: want %d, got %d",
				i+1, tc.s, tc.expected, actual)
		}
	}
}
