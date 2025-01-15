package lc2429

import "testing"

func TestMinimizeXor(t *testing.T) {
	testcases := []struct {
		num1     int
		num2     int
		expected int
	}{
		{3, 5, 3},  //5 has two set bits. Construct 3; then 3 xor 3 == 0.
		{1, 12, 3}, // 12 is 0b1100; two set bits. Construct 3; then 1 xor 3 == 2.
	}

	for i, tc := range testcases {
		actual := minimizeXor(tc.num1, tc.num2)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%d, %d) failed: want %d, got %d",
				i+1, tc.num1, tc.num2, tc.expected, actual)
		}
	}
}
