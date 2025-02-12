package lc2342

import "sort"

func getDigitSum(num int) int {
	ans := 0
	for num > 0 {
		ans += (num % 10)
		num /= 10
	}
	return ans
}

func maximumSum(nums []int) int {
	// each num in nums is at most 10^9. Thus, the maximum possible sum of its digits
	// is 9 * 9 == 81 for 999'999'999, i.e. 10^9 - 1.
	const MAX_SUM = 9 * 9
	ans := -1
	groups := make(map[int][]int) // `groups` maps sum-of-digits to its original number(s).

	for _, num := range nums {
		digitSum := getDigitSum(num)
		groups[digitSum] = append(groups[digitSum], num)
	}

	for _, group := range groups {
		if len(group) >= 2 { // only consider groups of at least two numbers.
			// Sort the group in reverse order. Yeah, this is prime Golang...
			sort.Sort(sort.Reverse(sort.IntSlice(group)))

			ans = max(ans, group[0]+group[1])
		}
	}
	return ans
}
