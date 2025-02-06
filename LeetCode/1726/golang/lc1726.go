package lc1726

func tupleSameProduct(nums []int) int {
	multiples := make(map[int]int)
	ans := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			mul := nums[i] * nums[j]
			ans += 8 * multiples[mul]
			multiples[mul] += 1
		}
	}
	return ans
}
