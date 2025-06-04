package lc3403

import "testing"

func Test_answerString(t *testing.T) {
	testcases := []struct {
		word       string
		numFriends int
		expected   string
	}{
		{"dbca", 2, "dbc"},
		{"gggg", 4, "g"},
	}

	for i, tc := range testcases {
		actual := answerString(tc.word, tc.numFriends)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%q, %d) failed: want %q, got %q",
				i+1, tc.word, tc.numFriends, tc.expected, actual)
		}
	}
}
