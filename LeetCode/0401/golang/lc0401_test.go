package lc0401

import (
	"slices"
	"testing"
)

func Test_readBinaryWatch(t *testing.T) {
	testcases := []struct {
		turnedOn int
		expected []string
	}{
		{1, []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
		{9, []string{}},
	}

	for i, tc := range testcases {
		actual := readBinaryWatch(tc.turnedOn)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("Testcase readBinaryWatch#%02d (%d) failed: want:\n  %v, got:\n  %v",
				i+1, tc.turnedOn, tc.expected, actual)
		}
	}
}
