package lc1866

import "testing"

func Test_rearrangeSticks(t *testing.T) {
	testcases := []struct {
		n        int
		k        int
		expected int
	}{
		{3, 2, 3},           // [1,3,2], [2,3,1], [2,1,3]
		{5, 5, 1},           // [1,2,3,4,5]
		{20, 11, 647427950}, // if it works on this example, it works well.
	}

	for i, tc := range testcases {
		actual := rearrangeSticks(tc.n, tc.k)
		if actual != tc.expected {
			t.Errorf("Testcase rearrangeSticks#%02d (n=%d, k=%d) failed: want %d, got %d",
				i+1, tc.n, tc.k, tc.expected, actual)
		}
	}
}
