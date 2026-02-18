package lc0693

import (
	"strconv"
	"testing"
)

func Test_hasAlternatingBits(t *testing.T) {
	testcases := []struct {
		input    int
		expected bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{3, false},
		{4, false},
		{5, true},   // 101 in binary.
		{7, false},  // 111 in binary.
		{11, false}, // 1011 in binary.
		{1398100, false},
		{1398101, true},
		{1398102, false},
		{2863311529, false},
		{2863311530, true},
		{2863311531, false},
	}

	for i, tc := range testcases {
		actual := hasAlternatingBits(tc.input)
		if actual != tc.expected {
			t.Errorf("Testcase hasAlternatingBits#%02d (%d | binary: %s) failed: want %t, got %t",
				i+1, tc.input, strconv.FormatInt(int64(tc.input), 2), tc.expected, actual)
		}
	}
}
