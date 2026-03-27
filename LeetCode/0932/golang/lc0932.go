package lc0932

// Runtime complexity: O(n * logn)
// Auxiliary space: O(n) (for holding the answer array)
// Subjective level: medium+; also not that fun to work this one out.
// Solved on: 2026-03-26
func beautifulArray(n int) []int {
	ans := make([]int, n)
	for i := range n {
		ans[i] = i + 1
	}

	partition := func(left int, right int, bitmask int) int {
		nextToSwap := left
		for c := left; c <= right; c++ {
			if ans[c]&bitmask > 0 {
				ans[c], ans[nextToSwap] = ans[nextToSwap], ans[c]
				nextToSwap += 1
			}
		}
		return nextToSwap - 1
	}

	var recursively_divide func(left int, right int, bitmask int)
	recursively_divide = func(left int, right int, bitmask int) {
		if left >= right {
			return
		}

		mid := partition(left, right, bitmask)

		recursively_divide(left, mid, bitmask<<1)
		recursively_divide(mid+1, right, bitmask<<1)
	}

	recursively_divide(0, n-1, 1)
	return ans
}
