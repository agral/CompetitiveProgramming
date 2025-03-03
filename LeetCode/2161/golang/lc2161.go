package lc2161

func pivotArray(nums []int, pivot int) []int {
	ans := make([]int, 0, len(nums))
	for _, num := range nums {
		if num < pivot {
			ans = append(ans, num)
		}
	}
	for _, num := range nums {
		if num == pivot {
			ans = append(ans, num)
		}
	}
	for _, num := range nums {
		if num > pivot {
			ans = append(ans, num)
		}
	}
	return ans
}
