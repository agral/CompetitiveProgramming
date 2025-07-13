package lc2410

import "testing"

func Test_matchPlayersAndTrainers(t *testing.T) {
	testcases := []struct {
		players  []int
		trainers []int
		expected int
	}{
		{[]int{4, 7, 9}, []int{8, 2, 5, 8}, 2},
		{[]int{1, 1, 1}, []int{10}, 1},
	}

	for i, tc := range testcases {
		actual := matchPlayersAndTrainers(tc.players, tc.trainers)
		if actual != tc.expected {
			t.Errorf("Testcase %02d (%v | %v) failed: want %d, got %d",
				i+1, tc.players, tc.trainers, tc.expected, actual)
		}
	}
}
