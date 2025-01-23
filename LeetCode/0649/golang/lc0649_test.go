package lc0649

import "testing"

func Test_predictPartyVictory(t *testing.T) {
	testcases := []struct {
		senate   string
		expected string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
		{"RRDD", "Radiant"},
		{"RDRD", "Radiant"},
		{"DRRD", "Dire"},
		{"DDRRR", "Dire"},
		{"DDRRRR", "Radiant"},
	}

	for i, tc := range testcases {
		actual := predictPartyVictory(tc.senate)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%s) failed: want %s, got %s", i+1, tc.senate, tc.expected, actual)
		}
	}
}
