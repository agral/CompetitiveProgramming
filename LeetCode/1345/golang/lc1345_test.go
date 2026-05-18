package lc1345

import "testing"

func Test_minJumps(t *testing.T) {
	testcases := []struct {
		arr      []int
		expected int
	}{
		{[]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, 3}, // idx 0 -> 4 -> 3 -> 9, three jumps.
		{[]int{7}, 0},                      // idx 0 is already the end of array; zero jumps required.
		{[]int{7, 6, 9, 6, 9, 6, 9, 7}, 1}, // idx 0 -> 7, one jump.
	}

	for i, tc := range testcases {
		actual := minJumps(tc.arr)
		if actual != tc.expected {
			t.Errorf("Testcase minJumps#%02d (%v) failed: want %d, got %d",
				i+1, tc.arr, tc.expected, actual)
		}
	}
}
