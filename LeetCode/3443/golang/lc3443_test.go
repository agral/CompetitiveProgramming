package lc3443

import "testing"

func Test_maxDistance(t *testing.T) {
	testcases := []struct {
		s        string
		k        int
		expected int
	}{
		{"NWSE", 1, 3},   // NWNE -> NWN results in distance of 3 from origin.
		{"NSWWEW", 3, 6}, // "NNWWWW"" results in distance of 6 from origin.
	}

	for i, tc := range testcases {
		actual := maxDistance(tc.s, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%q | %d) failed: want %d, got %d",
				i+1, tc.s, tc.k, tc.expected, actual)
		}
	}
}
