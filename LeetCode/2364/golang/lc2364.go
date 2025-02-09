package lc2364

func countBadPairs(nums []int) int64 {
	var ans int64 = 0

	// Holds nums[key] - key; then:
	// a)     count[nums[i] - i] holds the number of "good" pairs;
	// b) i - count[nums[i] - i] holds the number of "bad" pairs
	count := make(map[int]int)

	for idx, val := range nums {
		ans += int64(idx - count[val-idx])
		count[val-idx]++
	}
	return ans
}
