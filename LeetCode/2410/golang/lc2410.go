package lc2410

import "slices"

// Runtime complexity: O(sort) + O(n) == O(sort)
// Auxiliary space: O(sort) + O(n) == O(sort)
func matchPlayersAndTrainers(players []int, trainers []int) int {
	slices.Sort(players)
	slices.Sort(trainers)
	len_players := len(players)
	ans := 0
	for j := range trainers {
		if trainers[j] >= players[ans] {
			ans += 1
			if ans == len_players {
				return ans
			}
		}
	}
	// This one is for the case of no matches:
	return ans
}
