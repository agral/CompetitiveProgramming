package lc3946

import "math"

// Runtime complexity: O(n * (n+budget)) == O(n^2 + n*budget)
// Auxiliary space: O(budget)
// Subjective level: hard. It's a 0-1 knapsack with extras...
// Solved on: 2026-11-27

// Note: had to look up the 0-1 knapsack algorithm. Have not used that since high school times...

func maximumSaleItems(items [][]int, budget int) int {
	lowestPrice := math.MaxInt32
	// optimal[i]: max value that can be obtained using items whose
	// total value stays below or equal to i.
	// Note: with each available item used once at most.
	optimal := make([]int, budget+1)

	for _, item := range items {
		factor, price := item[0], item[1]

		// The total delta-value of this item (excluding itself):
		// the total count of items with factor that is a multiple of this item's factor
		// (the problem states: this item's factor divides that factor, which is somewhat more vague).
		// NOTE: this trivially includes one free copy of every item bought!
		// NOTE: we're obviously counting the COUNT (the NUMBER) of items, not their VALUE (!)
		bonusValue := 0
		for _, item := range items { // nested for, but can't think of any better solution. Is this optimal?
			if item[0]%factor == 0 {
				bonusValue += 1
			}
		}

		// 0-1 knapsack
		lowestPrice = min(lowestPrice, price)
		for b := budget; b >= price; b-- {
			optimal[b] = max(optimal[b], optimal[b-price]+bonusValue)
		}
	}

	// Finally use the knapsack result, but also try using the yet-unspent budget
	// (for buying the cheapest items):
	ans := 0
	for b := range budget + 1 { // WA #01: range 0 to budget, budget included.
		ans = max(ans, optimal[b]+(budget-b)/lowestPrice)
	}
	return ans
}
