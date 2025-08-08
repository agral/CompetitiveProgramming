package lc0808

import "testing"

func Test_soupServings(t *testing.T) {
	testcases := []struct {
		n        int
		expected float64
	}{
		{50, 0.625},
		{100, 0.71875},
	}

	for i, tc := range testcases {
		actual := soupServings(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %f, got %f", i+1, tc.n, tc.expected, actual)
		}
	}
}
