package lc0067

import "testing"

func Test_addBinary(t *testing.T) {
	testcases := []struct {
		a        string
		b        string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"1", "1111111", "10000000"},
		{"1111111", "1", "10000000"},
		{"1000", "1000", "10000"},
		{"10101", "1010", "11111"},
		{"10101", "1011", "100000"},
	}

	for i, tc := range testcases {
		actual := addBinary(tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("Testcase addBinary#%02d (%s + %s) failed: want %s, got %s",
				i+1, tc.a, tc.b, tc.expected, actual)
		}
	}
}
