package lc0190

import "testing"

func Test_reverseBits(t *testing.T) {
	testcases := []struct {
		n        int
		expected int
	}{
		{43261596, 964176192},    // Example 1
		{2147483644, 1073741822}, // Example 2

		{0x00000001, 0x80000000},
		{0x01020304, 0x20c04080},
	}

	for i, tc := range testcases {
		actual := reverseBits(tc.n)
		if actual != tc.expected {
			t.Errorf("Testcase reverseBits#%02d (%x) failed: want %x, got %x",
				i+1, tc.n, tc.expected, actual)
		}
	}
}
