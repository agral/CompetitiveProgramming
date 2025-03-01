package lc2460

func applyOperations(nums []int) []int {
	// Apply the operation:
	SZ := len(nums)
	for i := 0; i < SZ-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
			i += 1 // micro-optimization: skip past the newly created zero.
			// if the next number is zero too, this operation
			// changes nothing in this case. Otherwise it's not applied.
		}
	}

	// Move zeroes to the end of the array.
	// But don't swap in-place, instead let's create a new array and move the numbers there,
	// with implicit zeroes remaining at the end.
	ans := make([]int, SZ)
	k := 0
	for i := range SZ {
		if nums[i] != 0 {
			ans[k] = nums[i]
			k += 1
		}
	}
	return ans
}
