package lc1625

import "testing"

func Test_added(t *testing.T) {
	testcases := []struct {
		input    string
		a        int
		expected string
	}{
		{"3456", 5, "3951"},
		{"5525", 2, "5727"},
	}
	for i, tc := range testcases {
		actual := added(tc.input, tc.a)
		if actual != tc.expected {
			t.Errorf("Testcase add#%02d (%s | a=%d) failed: want %s, got %s",
				i+1, tc.input, tc.a, tc.expected, actual)
		}
	}
}

func Test_rotated(t *testing.T) {
	testcases := []struct {
		input    string
		b        int
		expected string
	}{
		{"3456", 1, "6345"},
		{"3456", 2, "5634"},
		{"3456", 3, "4563"},
	}
	for i, tc := range testcases {
		actual := rotated(tc.input, tc.b)
		if actual != tc.expected {
			t.Errorf("Testcase add#%02d (%s | b=%d) failed: want %s, got %s",
				i+1, tc.input, tc.b, tc.expected, actual)
		}
	}
}

func Test_findLexSmallestString(t *testing.T) {
	testcases := []struct {
		s        string
		a        int
		b        int
		expected string
	}{
		{"5525", 9, 2, "2050"},
		{"74", 5, 1, "24"},
		{"0011", 4, 2, "0011"},
	}

	for i, tc := range testcases {
		actual := findLexSmallestString(tc.s, tc.a, tc.b)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s | a=%d | b=%d) failed: want %s, got %s",
				i+1, tc.s, tc.a, tc.b, tc.expected, actual)
		}
	}
}
