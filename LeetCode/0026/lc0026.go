package lc0026

func removeDuplicates(nums []int) int {
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[left] {
			left++
			if left < right {
				nums[left] = nums[right]
			}
		}
	}
	return left + 1
}
