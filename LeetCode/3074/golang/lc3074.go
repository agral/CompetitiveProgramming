package lc3074

import (
	"slices"
)

// Runtime complexity: O(sort)
// Auxiliary space: O(sort)
// Subjective level: easy
// Solved on: 2025-12-24
func minimumBoxes(apples []int, boxes []int) int {
	// 1. calculate the total capacity required by all the apples:
	totalApples := 0
	for _, a := range apples {
		totalApples += a
	}

	// 2. Sort the boxes in descending order w.r.t. their capacity (biggest boxes first):
	slices.SortFunc(boxes, func(lhs int, rhs int) int {
		return rhs - lhs
	})

	// 3. Keep adding boxes one by one, calculating their total capacity.
	//    Stop when total capacity >= sum. Return the number of used boxes.
	totalBoxes := 0
	ans := 0
	for totalBoxes < totalApples {
		totalBoxes += boxes[ans]
		ans += 1
	}
	return ans
}
