package lc0904

import "testing"

func Test_totalFruit(t *testing.T) {
	testcases := []struct {
		fruits   []int
		expected int
	}{
		//{[]int{1, 2, 1}, 3},
		//{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
	}

	for i, tc := range testcases {
		actual := totalFruit(tc.fruits)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v) failed: want %d, got %d", i+1, tc.fruits, tc.expected, actual)
		}
	}
}
