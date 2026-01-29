package lc2976

import "math"

// Runtime complexity: O(26^3) + O(len(source)) + O(len(original)).
// Note: len(source) == len(target); len(original) == len(changed) == len(cost).
// Auxiliary space: O(26^2)
// Subjective level: medium+
// Solved on: 2026-01-29
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	const NUM_LETTERS byte = 'z' - 'a' + 1
	const IMPOSSIBLE int = math.MaxInt32

	// 1. distance[s][t] == the minimum cost to change byte('a'+s) to byte('a'+t). Initially "infinity".
	distance := make([][]int, NUM_LETTERS)
	for s := range NUM_LETTERS {
		distance[s] = make([]int, NUM_LETTERS)
		for t := range NUM_LETTERS {
			distance[s][t] = IMPOSSIBLE
		}
	}

	// 2. Fill in the actual distances from `original`, `changed` & `cost` slices:
	for i := range original {
		s := int(original[i] - 'a')
		t := int(changed[i] - 'a')
		distance[s][t] = min(distance[s][t], cost[i])
	}

	// 3. Calculate the least expensive "round-trip" paths from each char to each char.
	// That's O(NUM_LETTERS^3), not sure if this is optimal.
	for m := range NUM_LETTERS {
		for s := range NUM_LETTERS {
			if distance[s][m] != IMPOSSIBLE {
				for t := range NUM_LETTERS {
					if distance[m][t] != IMPOSSIBLE {
						distance[s][t] = min(distance[s][t], distance[s][m]+distance[m][t])
					}
				}
			}
		}
	}

	ans := int64(0)
	// 4. Finally calculate the answer - it's the sum of optimal costs for each substitution.
	// Note that switching from same letter to same letter costs 0 - not reflected in the `distance`,
	// handled separately.
	for i := range source {
		if source[i] != target[i] {
			s := int(source[i] - 'a')
			t := int(target[i] - 'a')

			// handle the case of an actually impossible-to-create transition:
			if distance[s][t] == IMPOSSIBLE {
				return -1
			}
			ans += int64(distance[s][t])
		}
	}

	return ans
}
