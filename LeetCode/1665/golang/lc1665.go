package lc1665

import "slices"

// Runtime complexity: O(sort) == O(n*logn)
// Auxiliary space: O(sort); other than that it's just two vars => O(1).
// Subjective level: medium
// Solved on: 2026-05-12
func minimumEffort(tasks [][]int) int {
	// task := []int{energyToPerform, energyToStart}.

	// 1. Sort the tasks by the delta-energy in descending order.
	// Delta-energy being the difference between the energy needed
	// to start the task and the energy actually needed to perfom it.
	slices.SortFunc(tasks, func(lhs []int, rhs []int) int {
		return (lhs[0] - lhs[1]) - (rhs[0] - rhs[1])
	})

	// 2. Calculate the minimum amount of energy reqired to perform all the tasks
	// in the order they're now sorted in.
	ans := 0
	lastDiff := 0
	for _, task := range tasks {
		if lastDiff < task[1] {
			// Increase the total starting energy - the current amount does not suffice
			// to start the current task.
			ans += task[1] - lastDiff
			lastDiff = task[1] - task[0]
		} else {
			// Just use up the energy from the last diff, it still is sufficient.
			lastDiff -= task[0]
		}
	}

	return ans
}
