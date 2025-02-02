package lc1752

func check(nums []int) bool {
	length := len(nums)
	numRotations := 0
	for i := 0; i < length; i++ {
		if nums[i] > nums[(i+1)%length] {
			numRotations++
			if numRotations >= 2 {
				return false
			}
		}
	}
	return true
}
