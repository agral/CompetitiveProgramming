package lc2126

import "slices"

// Runtime complexity: O(sort) == O(n * logn)
// Auxiliary space: O(sort)
// Subjective level: easy
// Solved on: 2026-05-31
func asteroidsDestroyed(mass int, asteroids []int) bool {
	slices.Sort(asteroids)
	for _, m := range asteroids {
		if mass >= m {
			mass += m
		} else {
			return false
		}
	}
	return true
}
