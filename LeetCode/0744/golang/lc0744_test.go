package lc0744

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	testcases := []struct {
		letters  []byte
		target   byte
		expected byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'x', 'x', 'y', 'y'}, 'z', 'x'},
	}

	for i, tc := range testcases {
		actual := nextGreatestLetter(tc.letters, tc.target)
		if actual != tc.expected {
			t.Errorf("Testcase nextGreatestLetter#%02d (%v | %s) failed: want %s, got %s",
				i+1, tc.letters, string(tc.target), string(tc.expected), string(actual))
		}
	}
}
