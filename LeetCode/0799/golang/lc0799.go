package lc0799

// Runtime complexity: O(query_row^2)
// Auxiliary space: O(query_row)
// Subjective level: medium
// Solved on: 2026-02-14
func champagneTower(poured int, query_row int, query_glass int) float64 {
	// Pour all the `poured` champagne directly into the top glass (0, 0):
	flow_previous_row := []float64{float64(poured)}
	for curr_row := 1; curr_row <= query_row; curr_row++ {
		// As the tower's rows are 0-indexed, there are (n+1) glasses in nth row:
		flow_curr_row := make([]float64, curr_row+1)

		// For each of the glasses in the previous row, calculate how much of the champagne
		// spills equally to both descendant glasses. Remember that one glass of champagne
		// remains in the "parent" (top) glass:
		for g := range curr_row {
			excess_flow := 0.5 * (flow_previous_row[g] - 1.0)
			if excess_flow > 0.0 {
				flow_curr_row[g] += excess_flow
				flow_curr_row[g+1] += excess_flow
			}
		}
		// At the end of each iteration, the "current" flow becomes the "previous" flow
		flow_previous_row = flow_curr_row
	}

	// Return what's at the queried row's glass; cap the result at at most 1.0 (i.e. glass being full).
	return min(flow_previous_row[query_glass], 1.0)
}
