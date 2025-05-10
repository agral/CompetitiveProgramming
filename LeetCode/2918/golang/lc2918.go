package lc2918

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func minSum(nums1 []int, nums2 []int) int64 {
	sum1 := int64(0)
	zeroes1 := int64(0)
	for _, num := range nums1 {
		if num == 0 {
			zeroes1++
		} else {
			sum1 += int64(num)
		}
	}

	sum2 := int64(0)
	zeroes2 := int64(0)
	for _, num := range nums2 {
		if num == 0 {
			zeroes2++
		} else {
			sum2 += int64(num)
		}
	}

	if zeroes1 == 0 && sum1 < sum2+zeroes2 {
		return -1
	}
	if zeroes2 == 0 && sum2 < sum1+zeroes1 {
		return -1
	}

	return max(sum1+zeroes1, sum2+zeroes2)
}
