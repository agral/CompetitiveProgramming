package lc2126

import "testing"

func Test_asteroidsDestroyed(t *testing.T) {
	testcases := []struct {
		mass      int
		asteroids []int
		expected  bool
	}{
		{10, []int{3, 9, 19, 5, 21}, true},
		{5, []int{4, 9, 23, 4}, false},
	}

	for i, tc := range testcases {
		actual := asteroidsDestroyed(tc.mass, tc.asteroids)
		if actual != tc.expected {
			t.Errorf("Testcase asteroidsDestroyed#%02d (mass=%d, asteroids=%v) failed: want %t, got %t",
				i+1, tc.mass, tc.asteroids, tc.expected, actual)
		}
	}
}
