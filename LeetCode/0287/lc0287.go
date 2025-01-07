package lc0287

// Given an array of integers nums containing n + 1 integers where each integer
// is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and using only constant extra space.
func findDuplicate(nums []int) int {
	tortoise := nums[nums[0]]
	hare := nums[nums[nums[0]]]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
	}
	tortoise = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}
	return tortoise
}
