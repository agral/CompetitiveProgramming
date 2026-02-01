package lc1239

import "math/bits"

// Runtime complexity: O(2^n) (can't think of any better runtime complexity, but maybe can be done better?)
// Auxiliary space: O(n)
// Subjective level: hard
// Solved on: 2026-02-01
func maxLength(arr []string) int {
	makeBitmask := func(s string) (uint64, bool) {
		bitmask := uint64(0)
		for _, ch := range s {
			i := int(ch - 'a')
			if (bitmask & (1 << i)) > 0 {
				return 0, false
			}
			bitmask |= (1 << i)
		}
		return bitmask, true
	}

	masks := []uint64{}
	for _, s := range arr {
		mask, isValid := makeBitmask(s)
		if isValid {
			masks = append(masks, mask)
		}
	}

	var dfs func(s int, used uint64) int
	dfs = func(s int, used uint64) int {
		value := bits.OnesCount64(used)
		for i := s; i < len(masks); i++ {
			if (used & masks[i]) == 0 {
				value = max(value, dfs(i+1, used|masks[i]))
			}
		}
		return value
	}

	return dfs(0, 0)
}
