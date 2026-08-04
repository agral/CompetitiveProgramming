package lc3731

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: easy
// Solved on: 2026-08-04
func findMissingElements(nums []int) []int {
	// 1. Find the beginning and the end of the range - min and max.
	// This is O(n).
	L := len(nums)
	minn := nums[0]
	maxx := nums[0]
	for i := 1; i < L; i++ {
		if nums[i] < minn {
			minn = nums[i]
		} else if nums[i] > maxx {
			maxx = nums[i]
		}
	}

	// 2. The length of the range is maxx - minn + 1. Reserve a slice of bools of this length,
	// then mark all the present numbers in that range (indexed at n - minn of course).
	// At this point we also know exactly how many numbers are missing: it's the length of the
	// range `R` minus the actual slice's length `L`.
	// As an optimization, if no numbers are missing (R == L), return an empty slice immediately.
	// This is O(n) too.
	R := maxx - minn + 1
	if R == L {
		return []int{}
	}
	isPresent := make([]bool, R)
	for _, num := range nums {
		isPresent[num-minn] = true
	}

	// Make the answer slice of size equal the missing numbers' count `R` - `L`,
	// then fill in the missing numbers. This is O(n) as well.
	ans := make([]int, R-L)
	ansIdx := 0
	for i := 0; ansIdx < R-L; i++ {
		if !isPresent[i] {
			ans[ansIdx] = minn + i
			ansIdx += 1
		}
	}
	return ans
}
