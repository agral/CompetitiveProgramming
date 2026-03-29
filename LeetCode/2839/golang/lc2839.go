package lc2839

// Runtime complexity: O(4) == O(1)
// Auxiliary space: O(1)
// Subjective level: easy. Also feels largely like a waste of time,
// but it's a daily one on 2026-03-29.
// Solved on: 2026-03-29
func canBeEqual(s1 string, s2 string) bool {
	// strings equal with no flipping:
	if s1[0] == s2[0] && s1[1] == s2[1] && s1[2] == s2[2] && s1[3] == s2[3] {
		return true
	}
	// strings equal with indices 0 and 2 swapped in s2:
	if s1[0] == s2[2] && s1[1] == s2[1] && s1[2] == s2[0] && s1[3] == s2[3] {
		return true
	}
	// strings equal with indices 0 and 2, 1 and 3 swapped in s2:
	if s1[0] == s2[2] && s1[1] == s2[3] && s1[2] == s2[0] && s1[3] == s2[1] {
		return true
	}
	// strings equal with indices 1 and 3 swapped in s2:
	if s1[0] == s2[0] && s1[1] == s2[3] && s1[2] == s2[2] && s1[3] == s2[1] {
		return true
	}

	return false
}
