package lc3736

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-11-03
func minMoves(nums []int) int {
	// This is essentially the same question as LC0453 that I solved two days ago.
	// But the solution contained there required TWO full swipes over the array:
	// one to find the minimum number, and second one to calculate the difference between
	// the minimum and the ith number. Here we could have used the same algorithm
	// (the only difference is that we're looking for a maximum, not a minimum number);
	// but it can be further improved to only require ONE full swipe over the array:
	// in one swipe, keep track of the maximum number, and sum up all the entries.
	// At the end of the swipe, knowing the maximum number and the total of the array
	// it is easily determined how many operations are needed to turn all the numbers
	// to the maximum value.

	L := len(nums)
	maxValue := nums[0]
	sum := nums[0]
	for i := 1; i < L; i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
		sum += nums[i]
	}

	return (L * maxValue) - sum
}
