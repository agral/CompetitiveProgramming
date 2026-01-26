package lc0858

// Runtime complexity: O(log(p))
// Auxiliary space: O(1)
// Subjective level: hard-. Hard to work these cases out, even with pen & paper.
// Solved on: 2026-01-26
func mirrorReflection(p int, q int) int {
	// 1. Simplify as much as possible:
	for p%2 == 0 && q%2 == 0 {
		p /= 2
		q /= 2
	}

	// 2. Now it's just a matter of whether p/q are evenly divisible by 2.
	if p%2 == 0 {
		return 2 // top left corner will be hit.
	}
	if q%2 == 0 {
		return 0 // bottom right corner will be hit.
	}
	return 1 // top right corner will be hit.
}
